  var socket = new WebSocket('ws://localhost:9000/ws')

let connect = (cb)=>{
    console.log("connecting")

    socket.onopen = ()=>{
        console.log("Webscoket connected successfully")
    }

    socket.onmessage = (msg)=>{
        console.log(msg)
        cb(msg)
    }

    socket.onClose =(event)=>{
        console.log("Webscoket closed",event)
    }


    socket.onerror = (error) =>{
        console.log("websocket error",error)
    }

}


let sendMsg = (msg) => {
    console.log("msg send" , msg)
    socket.send(msg)
}

export {connect,sendMsg}





