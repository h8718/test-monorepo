import { useNamespaceService } from "../../hooks";
import { useEffect, useState } from "react";
import { VscLayers, VscAdd } from "react-icons/vsc";
import { useNavigate, useParams } from "react-router";
import { Service } from ".";
import Alert from "../../components/alert";
import Button from "../../components/button";
import ContentPanel, {
  ContentPanelBody,
  ContentPanelTitle,
  ContentPanelTitleIcon,
  ContentPanelFooter,
} from "../../components/content-panel";
import FlexBox from "../../components/flexbox";
import Modal from "../../components/modal";
import { Config } from "../../util";
import Tippy from "@tippyjs/react";
import "tippy.js/dist/tippy.css";

import { useApiKey } from "../../util/apiKeyProvider";

export default function NamespaceRevisionsPanel(props) {
  const { namespace } = props;
  const { service } = useParams();

  if (!namespace) {
    return null;
  }

  return <NamespaceRevisions namespace={namespace} service={service} />;
}

export function RevisionCreatePanel(props) {
  const {
    image,
    setImage,
    scale,
    setScale,
    size,
    setSize,
    cmd,
    setCmd,
    traffic,
    setTraffic,
    maxScale,
  } = props;

  return (
    <FlexBox col gap style={{ fontSize: "12px" }}>
      <FlexBox col gap>
        <FlexBox col style={{ paddingRight: "10px" }}>
          Image
          <input
            value={image}
            onChange={(e) => setImage(e.target.value)}
            placeholder="Enter an image name"
          />
        </FlexBox>
        <FlexBox col style={{ paddingRight: "10px" }}>
          Scale
          <Tippy content={scale} trigger="mouseenter click">
            <input
              type="range"
              style={{ paddingLeft: "0px" }}
              min="0"
              max={maxScale.toString()}
              value={scale.toString()}
              onChange={(e) => setScale(e.target.value)}
            />
          </Tippy>
          <datalist
            style={{ display: "flex", alignItems: "center" }}
            id="sizeMarks"
          >
            <option
              style={{
                flex: "auto",
                textAlign: "left",
                lineHeight: "10px",
                paddingLeft: "8px",
              }}
              value="0"
              label="0"
            />
            <option
              style={{
                flex: "auto",
                textAlign: "right",
                lineHeight: "10px",
                paddingRight: "5px",
              }}
              value={maxScale}
              label={maxScale}
            />
          </datalist>
        </FlexBox>
        <FlexBox col style={{ paddingRight: "10px" }}>
          Size
          <input
            list="sizeMarks"
            style={{ paddingLeft: "0px" }}
            type="range"
            min="0"
            value={size.toString()}
            max="2"
            onChange={(e) => setSize(e.target.value)}
          />
          <datalist
            style={{ display: "flex", alignItems: "center" }}
            id="sizeMarks"
          >
            <option
              style={{ flex: "auto", textAlign: "left", lineHeight: "10px" }}
              value="0"
              label="small"
            />
            <option
              style={{ flex: "auto", textAlign: "center", lineHeight: "10px" }}
              value="1"
              label="medium"
            />
            <option
              style={{ flex: "auto", textAlign: "right", lineHeight: "10px" }}
              value="2"
              label="large"
            />
          </datalist>
        </FlexBox>
        <FlexBox col style={{ paddingRight: "10px" }}>
          CMD
          <input
            value={cmd}
            onChange={(e) => setCmd(e.target.value)}
            placeholder="Enter the CMD for a service"
          />
        </FlexBox>
        <FlexBox col style={{ paddingRight: "10px" }}>
          Traffic
          <Tippy content={`${traffic}%`} trigger="mouseenter click">
            <input
              type="range"
              style={{ paddingLeft: "0px" }}
              min="0"
              max="100"
              value={traffic.toString()}
              onChange={(e) => setTraffic(e.target.value)}
            />
          </Tippy>
          <datalist
            style={{ display: "flex", alignItems: "center" }}
            id="sizeMarks"
          >
            <option
              style={{ flex: "auto", textAlign: "left", lineHeight: "10px" }}
              value={0}
              label="0%"
            />
            <option
              style={{ flex: "auto", textAlign: "right", lineHeight: "10px" }}
              value={100}
              label="100%"
            />
          </datalist>
        </FlexBox>
      </FlexBox>
    </FlexBox>
  );
}

function NamespaceRevisions(props) {
  const { namespace, service } = props;
  const navigate = useNavigate();
  const [apiKey] = useApiKey();

  const {
    revisions,
    config,
    traffic,
    deleteNamespaceServiceRevision,
    getNamespaceServiceConfig,
    createNamespaceServiceRevision,
  } = useNamespaceService(Config.url, namespace, service, navigate, apiKey);

  const [load, setLoad] = useState(true);
  const [image, setImage] = useState("");
  const [scale, setScale] = useState(0);
  const [size, setSize] = useState(0);
  const [trafficPercent, setTrafficPercent] = useState(100);
  const [cmd, setCmd] = useState("");
  const [maxScale, setMaxScale] = useState(0);

  useEffect(() => {
    if (revisions !== null && revisions.length > 0) {
      setScale(revisions[0].minScale);
      setSize(revisions[0].size);
      setImage(revisions[0].image);
      setCmd(revisions[0].cmd);
    }
  }, [revisions]);

  useEffect(() => {
    async function cfgGet() {
      try {
        await getNamespaceServiceConfig().then((response) =>
          setMaxScale(response.maxscale)
        );
      } catch (e) {
        if (e.message === "get namespace service: not found") {
          navigate(`/not-found`);
        }
      }
    }
    if (load && config === null) {
      cfgGet();
      setLoad(false);
    }
  }, [config, getNamespaceServiceConfig, load, navigate]);

  if (revisions === null || traffic === null) {
    return null;
  }

  return (
    <FlexBox gap wrap style={{ paddingRight: "8px" }}>
      <FlexBox style={{ flex: 6 }}>
        <ContentPanel style={{ width: "100%" }}>
          <ContentPanelTitle>
            <ContentPanelTitleIcon>
              <VscLayers />
            </ContentPanelTitleIcon>
            <FlexBox>Service &apos;{service}&apos; Revisions</FlexBox>
            <div>
              <Modal
                title={`New '${service}' revision`}
                escapeToCancel
                modalStyle={{
                  maxWidth: "300px",
                }}
                button={<VscAdd />}
                buttonProps={{
                  auto: true,
                }}
                keyDownActions={[
                  {
                    code: "Enter",
                    fn: async () => {
                      return true;
                    },
                    closeModal: true,
                  },
                ]}
                requiredFields={[{ tip: "image is required", value: image }]}
                actionButtons={[
                  {
                    label: "Add",
                    onClick: async () => {
                      await createNamespaceServiceRevision(
                        image,
                        parseInt(scale),
                        parseInt(size),
                        cmd,
                        parseInt(trafficPercent)
                      );
                    },
                    buttonProps: { variant: "contained", color: "primary" },
                    closesModal: true,
                    validate: true,
                  },
                  {
                    label: "Cancel",
                    closesModal: true,
                  },
                ]}
              >
                {config !== null ? (
                  <RevisionCreatePanel
                    image={image}
                    setImage={setImage}
                    scale={scale}
                    setScale={setScale}
                    size={size}
                    setSize={setSize}
                    cmd={cmd}
                    setCmd={setCmd}
                    traffic={trafficPercent}
                    setTraffic={setTrafficPercent}
                    maxScale={maxScale}
                  />
                ) : (
                  ""
                )}
              </Modal>
            </div>
          </ContentPanelTitle>
          <ContentPanelBody>
            <FlexBox col gap>
              <FlexBox col gap>
                {revisions
                  .sort((a, b) => (a.created > b.created ? -1 : 1))
                  .map((obj, i) => {
                    let dontDelete = false;
                    let t = 0;
                    for (let i = 0; i < traffic.length; i++) {
                      if (traffic[i].revisionName === obj.name) {
                        dontDelete = true;
                        t = traffic[i].traffic;
                        break;
                      }
                    }
                    return (
                      <Service
                        key={obj.name}
                        latest={i === 0}
                        traffic={t}
                        dontDelete={dontDelete && i !== 0}
                        revision={obj.rev}
                        deleteService={deleteNamespaceServiceRevision}
                        url={`/n/${namespace}/services/${service}/${obj.rev}`}
                        conditions={obj.conditions}
                        name={obj.name}
                        status={obj.status}
                      />
                    );
                  })}
              </FlexBox>
            </FlexBox>
          </ContentPanelBody>
        </ContentPanel>
      </FlexBox>
      {/* <UpdateTraffic setNamespaceServiceRevisionTraffic={setNamespaceServiceRevisionTraffic} service={service} revisions={revisions} traffic={traffic}/> */}
    </FlexBox>
  );
}

export function UpdateTraffic(props) {
  const { traffic, service, revisions, setNamespaceServiceRevisionTraffic } =
    props;
  const [revOne, setRevOne] = useState(
    traffic[0] ? traffic[0].revisionName : ""
  );
  const [revTwo, setRevTwo] = useState(
    traffic[1] ? traffic[1].revisionname : ""
  );
  const [tpercent, setTPercent] = useState(traffic[0] ? traffic[0].traffic : 0);
  const [errMsg, setErrMsg] = useState("");

  // handle data from traffic stream updating
  useEffect(() => {
    if (traffic[0]) {
      setRevOne(traffic[0].revisionName);
      setTPercent(traffic[0].traffic);
    } else {
      setRevOne("");
      setTPercent(0);
    }
    if (traffic[1]) {
      setRevTwo(traffic[1].revisionName);
    } else {
      setRevTwo("");
    }
  }, [traffic]);

  return (
    <FlexBox style={{ flex: 1, minWidth: "370px" }}>
      <FlexBox gap style={{ fontSize: "12px", maxHeight: "fit-content" }}>
        <ContentPanel style={{ width: "100%", height: "fit-content" }}>
          <ContentPanelTitle>
            <ContentPanelTitleIcon>
              <VscLayers />
            </ContentPanelTitleIcon>
            <FlexBox>Update &apos;{service}&apos; traffic</FlexBox>
          </ContentPanelTitle>
          <ContentPanelBody className="secrets-panel">
            <FlexBox col gap style={{}}>
              <FlexBox col gap>
                <FlexBox col style={{ paddingRight: "4px" }}>
                  <span style={{ fontWeight: "bold" }}>Rev 1</span>
                  <select
                    value={revOne}
                    onChange={(e) => {
                      if (e.target.value === "") {
                        setTPercent(0);
                      }
                      setRevOne(e.target.value);
                    }}
                  >
                    <option value="">No revision selected</option>
                    {revisions.map((obj, key) => {
                      if (obj.name !== revTwo) {
                        return (
                          <option
                            key={`option-rev-update-traffic-1-${key}`}
                            value={obj.name}
                          >
                            {obj.name}
                          </option>
                        );
                      } else {
                        return null;
                      }
                    })}
                  </select>
                </FlexBox>
                <FlexBox col style={{ paddingRight: "4px" }}>
                  <span style={{ fontWeight: "bold" }}>Rev 2</span>
                  <select
                    value={revTwo}
                    onChange={(e) => {
                      if (e.target.value === "") {
                        setTPercent(100);
                      }
                      setRevTwo(e.target.value);
                    }}
                  >
                    <option value="">No revision selected</option>
                    {revisions.map((obj, key) => {
                      if (obj.name !== revOne) {
                        return (
                          <option
                            key={`option-rev-update-traffic-2-${key}`}
                            value={obj.name}
                          >
                            {obj.name}
                          </option>
                        );
                      } else {
                        return null;
                      }
                    })}
                  </select>
                </FlexBox>
                <FlexBox col style={{ paddingRight: "10px" }}>
                  <span style={{ fontWeight: "bold" }}>
                    Traffic Distribution
                  </span>
                  <FlexBox>
                    <FlexBox>Rev 1</FlexBox>
                    <FlexBox
                      style={{ textAlign: "right", justifyContent: "flex-end" }}
                    >
                      Rev 2
                    </FlexBox>
                  </FlexBox>
                  <input
                    disabled={revTwo === "" || revOne === "" ? true : false}
                    id="revisionMarks"
                    style={{ paddingLeft: "0px" }}
                    value={tpercent}
                    onChange={(e) => setTPercent(e.target.value)}
                    type="range"
                  />
                  <datalist
                    style={{ display: "flex", alignItems: "center" }}
                    id="revisionMarks"
                  >
                    <option
                      style={{
                        flex: "auto",
                        textAlign: "left",
                        lineHeight: "10px",
                      }}
                      value="0"
                      label={`${tpercent}%`}
                    />
                    <option
                      style={{
                        flex: "auto",
                        textAlign: "right",
                        lineHeight: "10px",
                      }}
                      value="100"
                      label={`${100 - tpercent}%`}
                    />
                  </datalist>
                </FlexBox>
                <FlexBox>
                  {errMsg ? (
                    <Alert
                      severity="error"
                      variant="filled"
                      grow
                      onClose={() => {
                        setErrMsg("");
                      }}
                    >
                      {errMsg}
                    </Alert>
                  ) : null}
                </FlexBox>
              </FlexBox>
            </FlexBox>
          </ContentPanelBody>
          <ContentPanelFooter>
            <FlexBox col style={{ alignItems: "flex-end" }}>
              <Button
                onClick={async () => {
                  try {
                    await setNamespaceServiceRevisionTraffic(
                      revOne,
                      parseInt(tpercent),
                      revTwo,
                      parseInt(100 - tpercent)
                    );
                    setErrMsg("");
                  } catch (err) {
                    if (err.message) {
                      setErrMsg(err.message);
                    } else {
                      setErrMsg(err);
                    }
                  }
                }}
              >
                Save
              </Button>
            </FlexBox>
          </ContentPanelFooter>
        </ContentPanel>
      </FlexBox>
    </FlexBox>
  );
}
