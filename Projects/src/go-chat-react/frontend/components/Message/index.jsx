import React, { useState } from 'react'

function index({Message}) {
    let temp = JSON.parse(Message)
    // const [message, setMessage] = useState();
    // setMessage(temp)
    console.log(temp)
  return (
    <div>
        {temp.body}
    </div>
  )
}

export default index