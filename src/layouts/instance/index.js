import React, { useEffect, useState } from 'react'
import './style.css'
import { Config, copyTextToClipboard } from '../../util';
import Button from '../../components/button';
import { useParams } from 'react-router';
import ContentPanel, { ContentPanelBody, ContentPanelFooter, ContentPanelHeaderButton, ContentPanelHeaderButtonIcon, ContentPanelTitle, ContentPanelTitleIcon } from '../../components/content-panel';
import FlexBox from '../../components/flexbox';
import {AiFillCode} from 'react-icons/ai';
import {useInstance, useInstanceLogs, useWorkflow} from 'direktiv-react-hooks';
import { CancelledState, FailState, RunningState, SuccessState } from '../instances';

import { Link } from 'react-router-dom';
import { AutoSizer, List, CellMeasurer, CellMeasurerCache } from 'react-virtualized';
import { IoCopy, IoEye, IoEyeOff } from 'react-icons/io5';
import * as dayjs from "dayjs"
import YAML from 'js-yaml'

import DirektivEditor from '../../components/editor';
import WorkflowDiagram from '../../components/diagram';
import { HiOutlineArrowsExpand } from 'react-icons/hi';
import Modal, { ButtonDefinition } from '../../components/modal';

function InstancePageWrapper(props) {

    let {namespace} = props;
    if (!namespace) {
        return <></>
    }

    return <InstancePage namespace={namespace} />

}

export default InstancePageWrapper;

function InstancePage(props) {

    let {namespace} = props;
    const [load, setLoad] = useState(true)
    const [wfpath, setWFPath] = useState("")
    const [rev, setRev] = useState("")
    const [follow, setFollow] = useState(true)
    const [width, setWidth] = useState(window.innerWidth);
    const [clipData, setClipData] = useState(null)
    const params = useParams()

    let instanceID = params["id"];

    let {data, err, cancelInstance, getInput, getOutput} = useInstance(Config.url, true, namespace, instanceID);
    

    useEffect(()=>{
        if(load && data !== null) {
            let split = data.as.split(":")
            setWFPath(split[0])
            setRev(split[1])
            setLoad(false)
        }
    },[load, data])

    if (data === null) {
        return <></>
    }

    if (err !== null) {
        // TODO
        return <></>
    }

    let label = <></>;
    if (data.status === "complete") {
        label = <SuccessState />
    } else if (data.status === "failed" || data.status === "crashed") {
        label = <FailState />
    }  else  if (data.status === "running") {
        label = <RunningState />
    } else if (data.status === "cancelled") {
        label = <CancelledState />
    }

    let wfName = data.as.split(":")[0]
    let revName = data.as.split(":")[1]

    let linkURL = "";
    if (!revName) {
        revName = "latest"
    }
    linkURL = `/n/${namespace}/explorer/${wfName}?tab=1&revision=${revName}&revtab=0`;

    return (<>
        <FlexBox className="col gap" style={{paddingRight: "8px"}}>
            <FlexBox className="gap wrap" style={{minHeight: "50%"}}>
                <FlexBox style={{minWidth: "340px", flex: "5"}}>
                    <ContentPanel style={{width: "100%"}}>
                        <ContentPanelTitle>
                            <ContentPanelTitleIcon>
                                <AiFillCode />
                            </ContentPanelTitleIcon>
                            <FlexBox className="gap" style={{alignItems:"center"}}>
                                <div>
                                    Instance Details
                                </div>
                                {label} 
                                <FlexBox style={{flex: "auto", justifyContent: "right", paddingRight: "6px", alignItems: "center"}}>
                                    <Link to={linkURL}>
                                        <Button className="small light">
                                            <span className="hide-on-small">View</span> Workflow
                                        </Button>
                                    </Link>
                                    <Modal
                                    escapeToCancel
                                    activeOverlay
                                    maximised
                                    noPadding
                                    title="Instance Details"
                                    titleIcon={
                                        <AiFillCode />
                                    }
                                    style={{
                                        maxWidth: "50px"
                                    }}
                                    modalStyle={{
                                        overflow: "hidden",
                                        padding: "0px"
                                    }}
                                    button={(
                                        <ContentPanelHeaderButton hackyStyle={{ marginBottom: "8px", height: "29px" }}>
                                            <ContentPanelHeaderButtonIcon>
                                                <HiOutlineArrowsExpand />
                                            </ContentPanelHeaderButtonIcon>
                                        </ContentPanelHeaderButton>
                                    )}
                                    actionButtons={[
                                        ButtonDefinition("Close", () => {}, "small light", true, false)
                                    ]}
                                >
                                    <InstanceLogs noPadding namespace={namespace} instanceID={instanceID} follow={follow} setFollow={setFollow} width={width} clipData={clipData} />
                                </Modal>
                                </FlexBox>
                            </FlexBox>
                        </ContentPanelTitle>
                        <InstanceLogs namespace={namespace} instanceID={instanceID} follow={follow} setFollow={setFollow} width={width} clipData={clipData} />
                    </ContentPanel>
                </FlexBox>
                <FlexBox className="gap wrap" style={{minWidth: "300px", flex: "2", flexWrap: "wrap-reverse"}}>
                    <FlexBox style={{minWidth: "300px"}}>
                        <ContentPanel style={{width: "100%"}}>
                        <ContentPanelTitle>
                            <ContentPanelTitleIcon>
                                <AiFillCode />
                            </ContentPanelTitleIcon>
                            <FlexBox className="gap">
                                <div>
                                Input
                                </div>
                            </FlexBox>
                            <Modal
                                escapeToCancel
                                activeOverlay                                
                                maximised
                                noPadding
                                title="Input"
                                titleIcon={
                                    <AiFillCode />
                                }
                                modalStyle={{
                                    overflow: "hidden",
                                    padding: "0px"
                                }}
                                button={(
                                    <ContentPanelHeaderButton>
                                        <ContentPanelHeaderButtonIcon>
                                            <HiOutlineArrowsExpand />
                                        </ContentPanelHeaderButtonIcon>
                                    </ContentPanelHeaderButton>
                                )}
                                actionButtons={[
                                    ButtonDefinition("Close", () => {}, "small light", true, false)
                                ]}
                            >
                                <Input getInput={getInput}/>
                            </Modal>
                        </ContentPanelTitle>
                        <ContentPanelBody>
                            <Input getInput={getInput}/>
                        </ContentPanelBody>
                    </ContentPanel>
                    </FlexBox>
                </FlexBox>
            </FlexBox>
            <FlexBox className="gap wrap">
                <FlexBox style={{minWidth: "300px", flex: "5"}}>
                    <ContentPanel style={{width: "100%"}}>
                        <ContentPanelTitle>
                            <ContentPanelTitleIcon>
                                <AiFillCode />
                            </ContentPanelTitleIcon>
                            <FlexBox className="gap" style={{alignItems:"center"}}>
                                <div>
                                    Logical Flow Graph
                                </div>
                            </FlexBox>
                        </ContentPanelTitle>
                        <ContentPanelBody>
                            <InstanceDiagram status={data.status} namespace={namespace} wfpath={wfpath} rev={rev} flow={data.flow}/>
                        </ContentPanelBody>
                    </ContentPanel>
                </FlexBox>
                <FlexBox style={{minWidth: "300px", flex: "2"}}>
                    <ContentPanel style={{width: "100%"}}>
                        <ContentPanelTitle>
                            <ContentPanelTitleIcon>
                                <AiFillCode />
                            </ContentPanelTitleIcon>
                            <FlexBox className="gap">
                                <div>
                                Output
                                </div>
                            </FlexBox>
                            <Modal
                                escapeToCancel
                                activeOverlay
                                maximised
                                noPadding
                                title="Output"
                                titleIcon={
                                    <AiFillCode />
                                }
                                modalStyle={{
                                    overflow: "hidden",
                                    padding: "0px"
                                }}
                                button={(
                                    <ContentPanelHeaderButton>
                                        <ContentPanelHeaderButtonIcon>
                                            <HiOutlineArrowsExpand />
                                        </ContentPanelHeaderButtonIcon>
                                    </ContentPanelHeaderButton>
                                )}
                                actionButtons={[
                                    ButtonDefinition("Close", () => {}, "small light", true, false)
                                ]}
                            >
                                <Output getOutput={getOutput} status={data.status}/>
                            </Modal>
                        </ContentPanelTitle>
                        <ContentPanelBody>
                            <Output getOutput={getOutput} status={data.status}/>
                        </ContentPanelBody>
                    </ContentPanel>
                </FlexBox>
            </FlexBox>
        </FlexBox>

    </>)
}

function InstanceLogs(props) {

    let {noPadding, namespace, instanceID, follow, setFollow, width, clipData} = props;

    let paddingStyle = { padding: "12px" }
    if (noPadding) {
        paddingStyle = { padding: "0px" }
    }

    return (
        <>
            <FlexBox className="col" style={{...paddingStyle}}>
                <FlexBox style={{ backgroundColor: "#002240", color: "white", borderRadius: "8px 8px 0px 0px", overflow: "hidden", padding: "8px" }}>
                    <Logs namespace={namespace} instanceID={instanceID} follow={follow} setFollow={setFollow} />
                </FlexBox>
                <div style={{ height: "40px", backgroundColor: "#223848", color: "white", maxHeight: "40px", minHeight: "40px", padding: "0px 10px 0px 10px", boxShadow: "0px 0px 3px 0px #fcfdfe", alignItems:'center', borderRadius: " 0px 0px 8px 8px", overflow: "hidden" }}>
                    <FlexBox className="gap" style={{width: "100%", flexDirection: "row-reverse", height: "100%", alignItems: "center"}}>
                        <TerminalButton onClick={()=>{
                            copyTextToClipboard(clipData)
                        }}>
                                <IoCopy/> Copy {width > 999 ? <span>to Clipboard</span>:""}
                        </TerminalButton>
                        {follow ?
                            <TerminalButton onClick={(e)=>setFollow(!follow)} className={"btn-terminal"}>
                                <IoEyeOff/> Stop {width > 999 ? <span>watching</span>: ""}
                            </TerminalButton>
                            :
                            <TerminalButton onClick={(e)=>setFollow(!follow)} className={"btn-terminal"} >
                                    <IoEye/> <div>Follow {width > 999 ? <span>logs</span>: ""}</div>
                            </TerminalButton>
                        }
                    </FlexBox>
                </div>
            </FlexBox>
        </>
    )
}

function TerminalButton(props) {

    let {children, onClick} = props;
    return (
        <div onClick={onClick} className="btn-terminal" style={{
            maxHeight: "22px"
        }}>
            <FlexBox className="gap" style={{ alignItems: "center", userSelect: "none" }}>
                {children}
            </FlexBox>
        </div>
    )

}

function InstanceDiagram(props) {
    const {wfpath, rev, flow, namespace, status} = props

    const [load, setLoad] = useState(true)
    const [wfdata, setWFData] = useState("")

    const {getWorkflowRevisionData} = useWorkflow(Config.url, false, namespace, wfpath)

    useEffect(()=>{
        async function getwf() {
            if(wfpath !== "" && rev !== "" && load){
                let ref = await getWorkflowRevisionData(rev)
                setWFData(atob(ref.revision.source))
                setLoad(false)
            }
        }
        
        getwf()
    },[wfpath, rev, load, getWorkflowRevisionData])

    if(load){
        return <></>
    }
    
    console.log(status, "STATUS OF INSTANCE")

    return(
        <WorkflowDiagram instanceStatus={status} disabled={true} flow={flow} workflow={YAML.load(wfdata)}/>
    )
}

function Input(props) {
    const {getInput} = props
    
    const [input, setInput] = useState("")

    useEffect(()=>{
        async function get() {
            if(input === "") {
                let data = await getInput()
                if(data === "") {
                    data = "No input data was provided..."
                }
                setInput(data)
            }
        }
        get()
    },[input, getInput])

    return(
        <FlexBox style={{overflow: "hidden"}}>
            {/* <div style={{width: "100%", height: "100%"}}> */}
            <AutoSizer>
                {({height, width})=>(
                    <DirektivEditor height={height} width={width} dlang="json" value={input} readonly={true}/>
                )}
            </AutoSizer>
            {/* </div> */}
        </FlexBox>
    )
}

function Output(props){
    const {getOutput, status} = props

    const [load, setLoad] = useState(true)
    const [output, setOutput] = useState("Waiting for instance to complete...")

    useEffect(()=>{
        async function get() {
            if (load && status !== "pending"){
                let data = await getOutput()
                if(data === "") {
                    data = "No output data was resolved..."
                } else {
                    data = JSON.stringify(JSON.parse(data), null, 2)
                }
                setOutput(data)
                setLoad(false)
            }
        }
        get()
    },[output, load, getOutput, status])

    useEffect(()=>{
        async function reGetOutput() {
            if(status !== "pending"){
                let data = await getOutput()
                if(data === "") {
                    data = "\"No output data was resolved...\""
                } else {
                    data = JSON.stringify(JSON.parse(data), null, 2)
                }
                setOutput(data)
            }
        }
       reGetOutput()
    },[status, getOutput])

    return(
        <FlexBox style={{padding: "0px", overflow: "hidden"}}>
            <AutoSizer>
                {({height, width})=>(
                    <DirektivEditor disableCursor height={height} width={width} dlang="json" value={output} readonly={true}/>
                )}
            </AutoSizer>
        </FlexBox>
    )
}

function InstanceTuple(props) {
    
    let {label, value, linkTo} = props;

    let x = value;
    if (linkTo) {
        x = (
            <Link to={linkTo}>{value}</Link>
        )
    }

    return (<>
        <FlexBox className="instance-details-tuple col" style={{minWidth: "150px", flex: "1"}}>
            <div>
                <b>{label}</b>
            </div>
            <div title={value} style={{fontSize: "12px", overflow: "hidden", textOverflow: "ellipsis", whiteSpace: "nowrap"}}>
                {x}
            </div>
        </FlexBox>
    </>)
}

function Logs(props){ 
    const cache = new CellMeasurerCache({
        fixedWidth: false,
        defaultHeight: 20
    })

    let {namespace, instanceID, follow} = props;
    // const [load, setLoad] = useState(true)
    // const [logs, setLogs] = useState([])
    let {data, err} = useInstanceLogs(Config.url, true, namespace, instanceID)
    
    // useEffect(()=>{
    //     if(load && data !== null){
    //         setLogs(data)
    //         setLoad(false)
    //     }
    // },[load])

    // useEffect(()=>{
    //     if(data !== null) {
    //         setLogs(logs + data)
    //     }
    // },[data, logs])


    if (!data) {
        return <></>
    }

    if (err) {
        return <></> // TODO 
    }

    function rowRenderer({index, parent, key, style}) {
        if(!data[index]){
            return ""
        }

        return (
        <CellMeasurer
            key={key}
            cache={cache}
            parent={parent}
            columnIndex={0}
            rowIndex={index}
        >
          <div style={style}>
            <div style={{display:"inline-block",minWidth:"112px", color:"#b5b5b5"}}>
                <div className="log-timestamp">
                    <div>[</div>
                        <div style={{display: "flex", flex: "auto", justifyContent: "center"}}>{dayjs.utc(data[index].node.t).local().format("HH:mm:ss.SSS")}</div>
                    <div>]</div>
                </div>
            </div> 
            <span style={{marginLeft:"5px"}}>
                {data[index].node.msg}
            </span>
            <div style={{height: `fit-content`}}></div>
          </div>
          </CellMeasurer>
        );
    }
      

    return (
        <WindowScroller>
            {({height, isScrolling, registerChild, scrollTop}) => {
                return (
                    <AutoSizer disableHeight>
                        {({ width }) => {
                            return (
                                <List 
                                    autoHeight
                                    height={height}
                                    isScrolling={isScrolling}
                                    rowCount={data.length}
                                    rowHeight={20}
                                    rowRenderer={rowRenderer}
                                    scrollTop={scrollTop}
                                    width={width}
                                />
                            )
                        }}
                    </AutoSizer>
                )
            }}
        </WindowScroller>
    )

    return(
        <div style={{flex:"1 1 auto", padding:'12px 12px 12px 12px', lineHeight: "20px"}}>
            <AutoSizer>
                {({height, width})=>(
                    <div style={{height: "100%", minHeight: "100%"}}>
                    <List
                    width={width}
                    height={height}
                        // style={{
                        //     minHeight: "100%"
                        //     // maxHeight: "100%"
                        // }}
                        rowRenderer={rowRenderer}
                        deferredMeasurementCache={cache}
                        scrollToIndex={follow ? data.length - 1: 0}
                        rowCount={data.length}
                        rowHeight={cache.rowHeight}
                        />
                    </div>
                )}
            </AutoSizer>
        </div>
    )
}