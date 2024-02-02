export default function Request_item( props ){
    return(
        <>
            <a href="/request_detail">
                <div id="request_name">{props.Title}</div>
                <div id="request_user_name">{props.Description}</div>
            </a>
        </>
    )
}