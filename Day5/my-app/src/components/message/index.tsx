import React from "react";
import "./index.css";

type MessageProps = {
    imgProfileUrl: string
    messageId: number
    senderName: string
    messageText: string
}

function Message(props: MessageProps): JSX.Element {
    return (
        <div className="message">
            <img className="profilePicture" src={props.imgProfileUrl} alt={props.senderName}/>
            <div>
                <p className="username">
                    {props.senderName}
                </p>
                <p className="messageContent">
                    {props.messageText}
                </p>
            </div>
        </div>
    );
}

export default Message