"use client"


import Image from "next/image";
import Head from "@/components/Head"
import ChatInput from "@/components/ChatInput"
import ChatHistory from "@/components/ChatHistory"
import { useEffect, useState } from "react";
import { sendMsg , connect} from "./api";

export default function Home() {

  const [chatHistory, setChatHistory] = useState([])
   
  useEffect(() => {
    connect((msg)=>{
      setChatHistory((prevChatHistory)=>[...prevChatHistory,msg]);
    })
    
  
    
  }, []);

 
  function send(event){
    sendMsg(event.target.value)
  }
  return (
    <main className="">
      <Head/>
      <ChatInput send={send}/>
      <ChatHistory chatHistory={chatHistory}/>
    </main>
  );
}
