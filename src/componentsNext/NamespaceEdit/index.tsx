import {
  DialogClose,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "~/design/Dialog";
import { GitCompare, Home, PlusCircle } from "lucide-react";
import {
  MirrorFormSchema,
  MirrorFormSchemaType,
  MirrorSshFormSchema,
  MirrorSshFormSchemaType,
  MirrorTokenFormSchema,
  MirrorTokenFormSchemaType,
  MirrorUpdateSshFormSchema,
  MirrorUpdateTokenFormSchema,
} from "~/api/namespaces/schema";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "~/design/Select";
import { SubmitHandler, useForm } from "react-hook-form";
import { Tabs, TabsList, TabsTrigger } from "~/design/Tabs";
import { useEffect, useState } from "react";

import Alert from "~/design/Alert";
import Button from "~/design/Button";
import FormErrors from "~/componentsNext/FormErrors";
import InfoTooltip from "./InfoTooltip";
import Input from "~/design/Input";
import { InputWithButton } from "~/design/InputWithButton";
import { MirrorInfoSchemaType } from "~/api/tree/schema/mirror";
import { Textarea } from "~/design/TextArea";
import { fileNameSchema } from "~/api/tree/schema/node";
import { pages } from "~/util/router/pages";
import { useCreateNamespace } from "~/api/namespaces/mutate/createNamespace";
import { useListNamespaces } from "~/api/namespaces/query/get";
import { useNamespaceActions } from "~/util/store/namespace";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { useUpdateMirror } from "~/api/tree/mutate/updateMirror";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

type FormInput = {
  name: string;
} & MirrorFormSchemaType &
  MirrorTokenFormSchemaType &
  MirrorSshFormSchemaType;

const mirrorAuthTypes = ["none", "ssh", "token"] as const;

type MirrorAuthType = (typeof mirrorAuthTypes)[number];

const NamespaceEdit = ({
  mirror,
  close,
}: {
  mirror?: MirrorInfoSchemaType;
  close: () => void;
}) => {
  const getInitialAuthType = (): MirrorAuthType => {
    if (!mirror) {
      return "none";
    }
    if (mirror?.info.url?.startsWith("git@")) {
      return "ssh";
    }
    if (mirror?.info.passphrase === "-") {
      return "token";
    }
    return "none";
  };

  const { t } = useTranslation();
  const [isMirror, setIsMirror] = useState(!!mirror);
  const [isNew] = useState(!mirror);
  const [authType, setAuthType] = useState<MirrorAuthType>(getInitialAuthType);
  const { data } = useListNamespaces();
  const { setNamespace } = useNamespaceActions();
  const navigate = useNavigate();

  const existingNamespaces = data?.results.map((n) => n.name) || [];

  const newNameSchema = fileNameSchema.and(
    z.string().refine((name) => !existingNamespaces.some((n) => n === name), {
      message: t("components.namespaceEdit.nameAlreadyExists"),
    })
  );

  // TODO: Improve this?
  // When a mirror is edited (in this case !isNew), the name isn't editable, so
  // no validation is needed for this value.
  const baseSchema = z.object({ name: isNew ? newNameSchema : z.string() });

  const getResolver = (
    isNew: boolean,
    isMirror: boolean,
    authType: MirrorAuthType
  ) => {
    if (!isMirror) {
      return zodResolver(baseSchema);
    }
    if (isNew && authType === "token") {
      return zodResolver(baseSchema.and(MirrorTokenFormSchema));
    }
    if (isNew && authType === "ssh") {
      return zodResolver(baseSchema.and(MirrorSshFormSchema));
    }
    if (!isNew && authType === "token") {
      return zodResolver(baseSchema.and(MirrorUpdateTokenFormSchema));
    }
    if (!isNew && authType === "ssh") {
      return zodResolver(baseSchema.and(MirrorUpdateSshFormSchema));
    }
    return zodResolver(baseSchema.and(MirrorFormSchema));
  };

  const {
    register,
    handleSubmit,
    trigger,
    formState: { isDirty, dirtyFields, errors, isValid, isSubmitted },
  } = useForm<FormInput>({
    resolver: getResolver(isNew, isMirror, authType),
    defaultValues: mirror
      ? {
          name: mirror.namespace,
          url: mirror.info.url,
          ref: mirror.info.ref,
        }
      : {},
  });

  const { mutate: createNamespace, isLoading } = useCreateNamespace({
    onSuccess: (data) => {
      setNamespace(data.namespace.name);
      navigate(
        pages.explorer.createHref({
          namespace: data.namespace.name,
        })
      );
      close();
    },
  });

  const { mutate: updateMirror } = useUpdateMirror({
    onSuccess: () => {
      close();
    },
  });

  const onSubmit: SubmitHandler<FormInput> = ({
    name,
    ref,
    url,
    passphrase,
    publicKey,
    privateKey,
  }) => {
    if (isNew) {
      createNamespace({
        name,
        mirror: { ref, url, passphrase, publicKey, privateKey },
      });
    }

    let updateAuthValues = {};

    if (authType === "none") {
      updateAuthValues = {
        passphrase: "",
        publicKey: "",
        privateKey: "",
      };
    }
    if (authType === "ssh") {
      const overwriteAuth =
        dirtyFields.passphrase ||
        dirtyFields.privateKey ||
        dirtyFields.publicKey;
      updateAuthValues = {
        passphrase: overwriteAuth ? passphrase : "-",
        publicKey: overwriteAuth ? publicKey : "-",
        privateKey: overwriteAuth ? privateKey : "-",
      };
    }
    if (authType === "token") {
      const overwriteAuth = dirtyFields.passphrase;
      updateAuthValues = {
        passphrase: overwriteAuth ? passphrase : "-",
        publicKey: "",
        privateKey: "",
      };
    }

    updateMirror({
      name,
      mirror: {
        ref,
        url,
        ...updateAuthValues,
      },
    });
  };

  // you can not submit if the form has not changed or if there are any errors and
  // you have already submitted the form (errors will first show up after submit)
  const disableSubmit = !isDirty || (isSubmitted && !isValid);

  // if the form has errors, we need to re-validate when isMirror or authType
  // has been changed, after useForm has updated the resolver.
  useEffect(() => {
    if (isSubmitted && !isValid) {
      trigger();
    }
  }, [isMirror, authType, isSubmitted, isValid, trigger]);

  const formId = `new-namespace`;
  return (
    <>
      <DialogHeader>
        <DialogTitle>
          {isNew ? <Home /> : <GitCompare />}
          {isNew
            ? t("components.namespaceEdit.title.new")
            : t("components.namespaceEdit.title.edit", {
                namespace: mirror?.namespace,
              })}
        </DialogTitle>
      </DialogHeader>

      {isNew && (
        <Tabs className="mt-2 sm:w-[400px]" defaultValue="namespace">
          <TabsList variant="boxed">
            <TabsTrigger
              variant="boxed"
              value="namespace"
              onClick={() => setIsMirror(false)}
            >
              {t("components.namespaceEdit.tab.namespace")}
            </TabsTrigger>
            <TabsTrigger
              variant="boxed"
              value="mirror"
              onClick={() => setIsMirror(true)}
            >
              {t("components.namespaceEdit.tab.mirror")}
            </TabsTrigger>
          </TabsList>
        </Tabs>
      )}

      <div className="mt-1 mb-3">
        <FormErrors errors={errors} className="mb-5" />
        <form
          id={formId}
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col gap-y-5"
        >
          {isNew && (
            <fieldset className="flex items-center gap-5">
              <label
                className="w-[112px] overflow-hidden text-right text-[14px]"
                htmlFor="name"
              >
                {t("components.namespaceEdit.label.name")}
              </label>
              <Input
                id="name"
                data-testid="new-namespace-name"
                placeholder={t("components.namespaceEdit.placeholder.name")}
                {...register("name")}
              />
            </fieldset>
          )}

          {isMirror && (
            <>
              <fieldset className="flex items-center gap-5">
                <label
                  className="w-[112px] flex-row overflow-hidden text-right text-[14px]"
                  htmlFor="url"
                >
                  {t("components.namespaceEdit.label.url")}
                </label>
                <InputWithButton>
                  <Input
                    id="url"
                    data-testid="new-namespace-url"
                    placeholder={t(
                      authType === "ssh"
                        ? "components.namespaceEdit.placeholder.gitUrl"
                        : "components.namespaceEdit.placeholder.httpUrl"
                    )}
                    {...register("url")}
                  />
                  <InfoTooltip>
                    {t("components.namespaceEdit.tooltip.url")}
                  </InfoTooltip>
                </InputWithButton>
              </fieldset>

              <fieldset className="flex items-center gap-5">
                <label
                  className="w-[112px] overflow-hidden text-right text-[14px]"
                  htmlFor="ref"
                >
                  {t("components.namespaceEdit.label.ref")}
                </label>
                <InputWithButton>
                  <Input
                    id="ref"
                    data-testid="new-namespace-ref"
                    placeholder={t("components.namespaceEdit.placeholder.ref")}
                    {...register("ref")}
                  />
                  <InfoTooltip>
                    {t("components.namespaceEdit.tooltip.ref")}
                  </InfoTooltip>
                </InputWithButton>
              </fieldset>

              <fieldset className="flex items-center gap-5">
                <label
                  className="w-[112px] overflow-hidden text-right text-[14px]"
                  htmlFor="ref"
                >
                  {t("components.namespaceEdit.label.authType")}
                </label>
                <Select
                  value={authType}
                  onValueChange={(value: MirrorAuthType) => setAuthType(value)}
                >
                  <SelectTrigger variant="outline" className="w-full">
                    <SelectValue
                      placeholder={t(
                        "components.namespaceEdit.placeholder.authType"
                      )}
                    />
                  </SelectTrigger>
                  <SelectContent>
                    {mirrorAuthTypes.map((option) => (
                      <SelectItem
                        key={option}
                        value={option}
                        onClick={() => setAuthType(option)}
                      >
                        {t(`components.namespaceEdit.authType.${option}`)}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </fieldset>

              {!isNew && authType !== "none" && (
                <Alert variant="info" className="text-sm">
                  {t("components.namespaceEdit.updateAuthInfo")}
                </Alert>
              )}

              {authType === "token" && (
                <fieldset className="flex items-center gap-5">
                  <label
                    className="w-[112px] overflow-hidden text-right text-[14px]"
                    htmlFor="token"
                  >
                    {t("components.namespaceEdit.label.token")}
                  </label>
                  <InputWithButton>
                    <Textarea
                      id="token"
                      data-testid="new-namespace-token"
                      placeholder={t(
                        "components.namespaceEdit.placeholder.token"
                      )}
                      {...register("passphrase")}
                    />
                    <InfoTooltip>
                      {t("components.namespaceEdit.tooltip.token")}
                    </InfoTooltip>
                  </InputWithButton>
                </fieldset>
              )}

              {authType === "ssh" && (
                <>
                  <fieldset className="flex items-center gap-5">
                    <label
                      className="w-[112px] overflow-hidden text-right text-[14px]"
                      htmlFor="passphrase"
                    >
                      {t("components.namespaceEdit.label.passphrase")}
                    </label>
                    <InputWithButton>
                      <Textarea
                        id="passphrase"
                        data-testid="new-namespace-passphrase"
                        placeholder={t(
                          "components.namespaceEdit.placeholder.passphrase"
                        )}
                        {...register("passphrase")}
                      />
                      <InfoTooltip>
                        {t("components.namespaceEdit.tooltip.passphrase")}
                      </InfoTooltip>
                    </InputWithButton>
                  </fieldset>
                  <fieldset className="flex items-center gap-5">
                    <label
                      className="w-[112px] overflow-hidden text-right text-[14px]"
                      htmlFor="public-key"
                    >
                      {t("components.namespaceEdit.label.publicKey")}
                    </label>
                    <InputWithButton>
                      <Textarea
                        id="public-key"
                        data-testid="new-namespace-pubkey"
                        placeholder={t(
                          "components.namespaceEdit.placeholder.publicKey"
                        )}
                        {...register("publicKey")}
                      />
                      <InfoTooltip>
                        {t("components.namespaceEdit.tooltip.publicKey")}
                      </InfoTooltip>
                    </InputWithButton>
                  </fieldset>

                  <fieldset className="flex items-center gap-5">
                    <label
                      className="w-[112px] overflow-hidden text-right text-[14px]"
                      htmlFor="private-key"
                    >
                      {t("components.namespaceEdit.label.privateKey")}
                    </label>
                    <InputWithButton>
                      <Textarea
                        id="private-key"
                        data-testid="new-namespace-privkey"
                        placeholder={t(
                          "components.namespaceEdit.placeholder.privateKey"
                        )}
                        {...register("privateKey")}
                      />
                      <InfoTooltip>
                        {t("components.namespaceEdit.tooltip.privateKey")}
                      </InfoTooltip>
                    </InputWithButton>
                  </fieldset>
                </>
              )}
            </>
          )}
        </form>
      </div>

      <DialogFooter>
        <DialogClose asChild>
          <Button variant="ghost">
            {t("components.namespaceEdit.cancelBtn")}
          </Button>
        </DialogClose>
        <Button
          data-testid="new-namespace-submit"
          type="submit"
          disabled={disableSubmit}
          loading={isLoading}
          form={formId}
        >
          {!isLoading && <PlusCircle />}
          {t("components.namespaceEdit.createBtn")}
        </Button>
      </DialogFooter>
    </>
  );
};

export default NamespaceEdit;
