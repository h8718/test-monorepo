import { useEffect, useState } from "react"
import { IoPlay } from "react-icons/io5"
import { useParams } from "react-router"
import { Service } from "../namespace-services"
import { RevisionCreatePanel, UpdateTraffic } from "../namespace-services/revisions"
import AddValueButton from "../../components/add-button"
import ContentPanel, { ContentPanelBody, ContentPanelTitle, ContentPanelTitleIcon } from "../../components/content-panel"
import FlexBox from "../../components/flexbox"
import Modal, { ButtonDefinition, KeyDownDefinition } from "../../components/modal"
import { Config } from "../../util"
import { useGlobalService } from "direktiv-react-hooks"

export default function GlobalRevisionsPanel(props){
    const {service} = useParams()
    const {revisions, config, traffic, createGlobalServiceRevision, deleteGlobalServiceRevision, setGlobalServiceRevisionTraffic, getServiceConfig} = useGlobalService(Config.url, service, localStorage.getItem("apikey"))

    const [load, setLoad] = useState(true)
    const [image, setImage] = useState("")
    const [scale, setScale] = useState(0)
    const [size, setSize] = useState(0)
    const [trafficPercent, setTrafficPercent] = useState(100)
    const [cmd, setCmd] = useState("")


    useEffect(()=>{
        if(revisions !== null) {
            setScale(revisions[0].minScale)
            setSize(revisions[0].size)
            setImage(revisions[0].image)
            setCmd(revisions[0].cmd)
        }
    },[revisions])

    useEffect(()=>{
        async function cfgGet() {
            await getServiceConfig()
        }
        if(load && config === null) {
            cfgGet()
            setLoad(false)
        }
    },[config, getServiceConfig, load])

    if(revisions === null) {
        return ""
    }


    return(
        <FlexBox className="gap wrap" style={{paddingRight:"8px"}}>
            <FlexBox  className="gap">
                    <FlexBox>
                        <ContentPanel style={{width:"100%"}}>
                        <ContentPanelTitle>
                            <ContentPanelTitleIcon>
                                <IoPlay/>
                            </ContentPanelTitleIcon>
                            <FlexBox>
                                Service '{service}' Revisions
                            </FlexBox>
                            <div>
                            <Modal title={`New '${service}' revision`} 
                                escapeToCancel
                                modalStyle={{
                                    maxWidth: "300px"
                                }}
                                onOpen={() => {
                                    console.log("ON OPEN");
                                }}
                                onClose={()=>{
                                }}
                                button={(
                                    <AddValueButton  label=" " />
                                )}  
                                keyDownActions={[
                                    KeyDownDefinition("Enter", async () => {
                                    }, true)
                                ]}
                                actionButtons={[
                                    ButtonDefinition("Add", async () => {
                                        let err = await createGlobalServiceRevision(image, parseInt(scale), parseInt(size), cmd, parseInt(trafficPercent))
                                        if (err) return err
                                    }, "small blue", true, false),
                                    ButtonDefinition("Cancel", () => {
                                    }, "small light", true, false)
                                ]}
                            >
                                {config !== null ? 
                                <RevisionCreatePanel 
                                    image={image} setImage={setImage}
                                    scale={scale} setScale={setScale}
                                    size={size} setSize={setSize}
                                    cmd={cmd} setCmd={setCmd}
                                    traffic={trafficPercent} setTraffic={setTrafficPercent}
                                    maxscale={config.maxscale}
                                />:""}
                            </Modal>
                        </div>
                        </ContentPanelTitle>
                            <ContentPanelBody className="secrets-panel">
                                <FlexBox className="gap col">
                                    <FlexBox className="col gap">
                                        {revisions.map((obj)=>{
                                            let dontDelete = false
                                            for(var i=0; i < traffic.length; i++) {
                                                if(traffic[i].revisionName === obj.name){
                                                    dontDelete= true
                                                    break
                                                }
                                            }
                                            return(
                                                <Service 
                                                    dontDelete={dontDelete}
                                                    revision={obj.rev}
                                                    deleteService={deleteGlobalServiceRevision}
                                                    url={`/g/services/${service}/${obj.rev}`}
                                                    conditions={obj.conditions}
                                                    name={obj.name}
                                                    status={obj.status}
                                                />
                                            )
                                        })}
                                    </FlexBox>
                                </FlexBox>
                            </ContentPanelBody>
                        </ContentPanel>
                    </FlexBox>
                    <UpdateTraffic setNamespaceServiceRevisionTraffic={setGlobalServiceRevisionTraffic} service={service} revisions={revisions} traffic={traffic}/>
                    </FlexBox>
        </FlexBox>
    )
}