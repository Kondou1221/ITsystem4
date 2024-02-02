import Request_item from "@/features/request/components/request_item"

export default function Request_list( props ){
    const list = []

    for (let [i, value] of Object.entries(props)) {
        console.log(props)
        list.push(<div id="request_item" key={i}><Request_item {...props[i]} /></div>)
    }    

    return (
        <div id="main_content">
            <div id="request_show">
                {list}
            </div>
        </div>
    )
}