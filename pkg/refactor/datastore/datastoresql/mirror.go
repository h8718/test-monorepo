package datastoresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/direktiv/direktiv/pkg/refactor/core"
	"github.com/direktiv/direktiv/pkg/refactor/mirror"
	"github.com/direktiv/direktiv/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sqlMirrorStore struct {
	// database connection.
	db *gorm.DB
	// symmetric encryption key to encrypt and decrypt mirror data.
	configEncryptionKey string
}

// SetVariable sets runtime variable into the database.
// nolint
func (s sqlMirrorStore) SetVariable(ctx context.Context, variable *core.RuntimeVariable) error {
	linkName := "namespace_id"
	linkValue := variable.NamespaceID

	if variable.WorkflowID.String() != (uuid.UUID{}).String() {
		linkName = "workflow_id"
		linkValue = variable.WorkflowID
	}

	res := s.db.WithContext(ctx).Exec(fmt.Sprintf(
		`UPDATE runtime_variables SET
						mime_type=?,
						data=?
					WHERE %s = ? AND name = ?;`, linkName),
		variable.MimeType, variable.Data, linkValue, variable.Name)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected > 1 {
		return fmt.Errorf("unexpected gorm update count, got: %d, want: %d", res.RowsAffected, 1)
	}
	if res.RowsAffected == 1 {
		return nil
	}

	newUUID := uuid.New()
	res = s.db.WithContext(ctx).Exec(fmt.Sprintf(`
							INSERT INTO runtime_variables(
								id, %s, name, mime_type, data) 
							VALUES(?, ?, ?, ?, ?);`, linkName),
		newUUID, linkValue, variable.Name, variable.MimeType, variable.Data)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return fmt.Errorf("unexpected runtime_variables insert count, got: %d, want: %d", res.RowsAffected, 1)
	}

	return nil
}

func cryptDecryptConfig(config *mirror.Config, key string, encrypt bool) (*mirror.Config, error) {
	resultConfig := &mirror.Config{}

	*resultConfig = *config

	targets := []*string{
		&resultConfig.PrivateKeyPassphrase,
		&resultConfig.PrivateKey,
	}

	for i := range targets {
		t := targets[i]

		var (
			b   string
			err error
		)
		if encrypt {
			b, err = util.EncryptDataBase64([]byte(key), []byte(*t))
		} else if len(*t) > 0 {
			b, err = util.DecryptDataBase64([]byte(key), *t)
		}

		if err != nil {
			return nil, err
		}
		*t = b
	}

	return resultConfig, nil
}

func (s sqlMirrorStore) CreateConfig(ctx context.Context, config *mirror.Config) (*mirror.Config, error) {
	newConfig, err := cryptDecryptConfig(config, s.configEncryptionKey, true)
	if err != nil {
		return nil, err
	}

	res := s.db.WithContext(ctx).Table("mirror_configs").Create(&newConfig)
	if res.Error != nil {
		return nil, res.Error
	}

	return s.GetConfig(ctx, newConfig.NamespaceID)
}

func (s sqlMirrorStore) UpdateConfig(ctx context.Context, config *mirror.Config) (*mirror.Config, error) {
	config, err := cryptDecryptConfig(config, s.configEncryptionKey, true)
	if err != nil {
		return nil, err
	}

	res := s.db.WithContext(ctx).Table("mirror_configs").
		Where("namespace_id", config.NamespaceID).
		Updates(config)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, fmt.Errorf("unexpected gorm update count, got: %d, want: %d", res.RowsAffected, 1)
	}

	return s.GetConfig(ctx, config.NamespaceID)
}

func (s sqlMirrorStore) GetConfig(ctx context.Context, namespaceID uuid.UUID) (*mirror.Config, error) {
	config := &mirror.Config{}
	res := s.db.WithContext(ctx).Raw(`
					SELECT *
					FROM mirror_configs
					WHERE namespace_id=?`, namespaceID).
		First(config)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, mirror.ErrNotFound
	}
	if res.Error != nil {
		return nil, res.Error
	}

	config, err := cryptDecryptConfig(config, s.configEncryptionKey, false)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (s sqlMirrorStore) CreateProcess(ctx context.Context, process *mirror.Process) (*mirror.Process, error) {
	newProcess := *process
	res := s.db.WithContext(ctx).Table("mirror_processes").Create(&newProcess)
	if res.Error != nil {
		return nil, res.Error
	}

	return &newProcess, nil
}

func (s sqlMirrorStore) UpdateProcess(ctx context.Context, process *mirror.Process) (*mirror.Process, error) {
	res := s.db.WithContext(ctx).Table("mirror_processes").
		Where("id", process.ID).
		Updates(process)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, fmt.Errorf("unexpected gorm update count, got: %d, want: %d", res.RowsAffected, 1)
	}

	return s.GetProcess(ctx, process.ID)
}

func (s sqlMirrorStore) GetProcess(ctx context.Context, id uuid.UUID) (*mirror.Process, error) {
	process := &mirror.Process{}
	res := s.db.WithContext(ctx).Raw(`
					SELECT *
					FROM mirror_processes
					WHERE id=?`, id).
		First(process)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, mirror.ErrNotFound
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return process, nil
}

func (s sqlMirrorStore) GetProcessesByNamespaceID(ctx context.Context, namespaceID uuid.UUID) ([]*mirror.Process, error) {
	var process []*mirror.Process

	res := s.db.WithContext(ctx).Raw(`
					SELECT *
					FROM mirror_processes
					WHERE namespace_id=?`, namespaceID).
		Find(&process)

	if res.Error != nil {
		return nil, res.Error
	}

	return process, nil
}

var _ mirror.Store = sqlMirrorStore{}
