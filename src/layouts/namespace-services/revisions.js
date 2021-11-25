import { useNamespaceService } from "direktiv-react-hooks"
import { useEffect, useState } from "react"
import { IoPlay } from "react-icons/io5"
import { useParams } from "react-router"
import { Service } from "."
import AddValueButton from "../../components/add-button"
import Button from "../../components/button"
import ContentPanel, { ContentPanelBody, ContentPanelTitle, ContentPanelTitleIcon, ContentPanelFooter } from "../../components/content-panel"
import FlexBox from "../../components/flexbox"
import Modal, { ButtonDefinition, KeyDownDefinition } from "../../components/modal"
import { Config } from "../../util"

export default function NamespaceRevisionsPanel(props) {
    const {namespace} = props
    const {service} = useParams()

    if(!namespace) {
        return ""
    }

    return(
        <FlexBox className="gap wrap" style={{paddingRight:"8px"}}>
            <NamespaceRevisions namespace={namespace} service={service} />
          
        </FlexBox>
    )
}

function RevisionCreatePanel(props){
    const {image, setImage, scale, setScale, size, setSize, cmd, setCmd, traffic, setTraffic, maxscale} = props

    return(
        <FlexBox className="col gap" style={{fontSize: "12px"}}>
            <FlexBox className="col gap">
                    <FlexBox className="col" style={{paddingRight:"10px"}}>
                        Image
                        <input value={image} onChange={(e)=>setImage(e.target.value)} placeholder="Enter an image name" />
                    </FlexBox>
                    <FlexBox className="col" style={{paddingRight:"10px"}}>
                        Scale
                        <input type="range" style={{paddingLeft:"0px"}} min={"0"} max={maxscale.toString()} value={scale.toString()} onChange={(e)=>setScale(e.target.value)} />
                    </FlexBox>
                    <FlexBox className="col" style={{paddingRight:"10px"}}>
                        Size
                        <input list="sizeMarks" style={{paddingLeft:"0px"}} type="range" min={"0"} value={size.toString()}  max={"2"} onChange={(e)=>setSize(e.target.value)}/>
                        <datalist style={{display:"flex", alignItems:'center'}} id="sizeMarks">
                            <option style={{flex:"auto", textAlign:"left", lineHeight:"10px"}} value="0" label="small"/>
                            <option style={{flex:"auto", textAlign:"center" , lineHeight:"10px"}} value="1" label="medium"/>
                            <option style={{flex:"auto", textAlign:"right", lineHeight:"10px" }} value="2" label="large"/>
                        </datalist>
                    </FlexBox>
                    <FlexBox className="col" style={{paddingRight:"10px"}}>
                        CMD
                        <input value={cmd} onChange={(e)=>setCmd(e.target.value)} placeholder="Enter the CMD for a service" />
                    </FlexBox>
                    <FlexBox className="col" style={{paddingRight:"10px"}}>
                        Traffic
                        <input type="range" style={{paddingLeft:"0px"}} min={"0"} max="100" value={traffic.toString()} onChange={(e)=>setTraffic(e.target.value)} />
                    </FlexBox>
            </FlexBox>
        </FlexBox>
    )
}

function NamespaceRevisions(props) {
    const {namespace, service} = props

    const {revisions, fn, config, traffic, err, setNamespaceServiceRevisionTraffic, deleteNamespaceServiceRevision, getNamespaceServiceConfig, createNamespaceServiceRevision} = useNamespaceService(Config.url, namespace, service)

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
            await getNamespaceServiceConfig()
        }
        if(load && config === null) {
            cfgGet()
            setLoad(false)
        }
    },[config, getNamespaceServiceConfig, load])

    if(revisions === null || traffic === null) {
        return ""
    }

    return(
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
                            let err = await createNamespaceServiceRevision(image, parseInt(scale), parseInt(size), cmd, parseInt(trafficPercent))
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
                                        deleteService={deleteNamespaceServiceRevision}
                                        url={`/n/${namespace}/services/${service}/${obj.rev}`}
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
        <UpdateTraffic setNamespaceServiceRevisionTraffic={setNamespaceServiceRevisionTraffic} service={service} revisions={revisions} traffic={traffic}/>
        </FlexBox>
    )
}

function UpdateTraffic(props){

    const {traffic, service, revisions, setNamespaceServiceRevisionTraffic} = props

    const [revOne, setRevOne] = useState(traffic[0] ? traffic[0].revisionName : "")
    const [revTwo, setRevTwo] = useState(traffic[1] ? traffic[1].revisionname : "")
    const [tpercent, setTPercent] = useState(traffic[0] ? traffic[0].traffic : 0)

    return(
        <FlexBox className="gap" style={{maxWidth:"300px", fontSize:"12px"}}>
            <ContentPanel style={{width:"100%", height:"fit-content"}}>
                <ContentPanelTitle>
                    <ContentPanelTitleIcon>
                        <IoPlay/>
                    </ContentPanelTitleIcon>
                    <FlexBox>
                        Update '{service}' traffic
                    </FlexBox>
                </ContentPanelTitle>
                    <ContentPanelBody className="secrets-panel">
                        <FlexBox className="gap col" style={{paddingLeft:"10px"}}>
                            <FlexBox className="col gap">
                                <FlexBox className="col" style={{paddingRight:"13px"}}>
                                    <span style={{fontWeight:"bold"}}>Rev 1</span>
                                    <select value={revOne} onChange={(e)=>{
                                        if(e.target.value === "") {
                                            setTPercent(0)
                                        }
                                        setRevOne(e.target.value)
                                    }}>
                                        <option value="">No revision selected</option>
                                        {revisions.map((obj)=>{
                                            if(obj.name !== revTwo) {
                                                return(
                                                    <option value={obj.name}>{obj.name}</option>
                                                )
                                            }
                                        })}
                                    </select>
                                </FlexBox>
                                <FlexBox className="col" style={{paddingRight:"13px"}}>
                                    <span style={{fontWeight:"bold"}}>Rev 2</span>
                                    <select value={revTwo} onChange={(e)=>{
                                        if(e.target.value === "") {
                                            setTPercent(100)
                                        }
                                        setRevTwo(e.target.value)
                                    }}>
                                        <option value="">No revision selected</option>
                                        {revisions.map((obj)=>{
                                            if(obj.name !== revOne) {
                                                return(
                                                    <option value={obj.name}>{obj.name}</option>
                                                )
                                            } 
                                        })}
                                    </select>
                                </FlexBox>
                                <FlexBox className="col" style={{paddingRight:"23px"}}>
                                    <span style={{fontWeight:"bold"}}>Traffic Distribution</span>
                                    <FlexBox>
                                        <FlexBox>
                                            Rev 1
                                        </FlexBox>
                                        <FlexBox style={{textAlign:"right", justifyContent:"flex-end"}}>
                                            Rev 2
                                        </FlexBox>
                                    </FlexBox>
                                    <input 
                                        disabled={revTwo === "" || revOne === "" ? true:false} 
                                        id="revisionMarks" 
                                        style={{paddingLeft:"0px"}} 
                                        value={tpercent} onChange={(e)=>setTPercent(e.target.value)} 
                                        type="range" 
                                    />
                                    <datalist style={{display:"flex", alignItems:'center'}} id="revisionMarks">
                                        <option style={{flex:"auto", textAlign:"left", lineHeight:"10px"}} value="0" label={`${tpercent}%`}/>
                                        <option style={{flex:"auto", textAlign:"right", lineHeight:"10px" }} value="100" label={`${100-tpercent}%`}/>
                                    </datalist>
                                </FlexBox>
                            </FlexBox>
                        </FlexBox>
                    </ContentPanelBody>
                    <ContentPanelFooter>
                        <FlexBox className="col" style={{paddingRight:"13px", alignItems:"flex-end"}}>
                            <Button className="small" onClick={async ()=>{
                                let err = await setNamespaceServiceRevisionTraffic(revOne, parseInt(tpercent), revTwo, parseInt(100-tpercent))
                                console.log(err, "not to sure how to display this err yet")

                            }}>
                                Save
                            </Button>
                        </FlexBox>
                    </ContentPanelFooter>
            </ContentPanel>
        </FlexBox>
    )
}