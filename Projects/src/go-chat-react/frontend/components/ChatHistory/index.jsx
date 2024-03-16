import React from 'react'
import Message from "@/components/Message"

function index({chatHistory}) {
  return (
    <div>
        {
            chatHistory.map(msg=>(<Message key={msg.timeStamp} Message={msg.data} />))
        }
    </div>
  )
}

export default index