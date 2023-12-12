import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Typography from '@mui/material/Typography';
import Paper from '@mui/material/Paper';
import {useState} from "react";
import Button from "@mui/material/Button";
import index from '../css/index.modules.css'

export default function ContentPage() {
        const [messages, setMessages] = useState([]);
        const [inputValue, setInputValue] = useState('');

        const handleInputChange = (e) => {
            setInputValue(e.target.value);
        };

        const handleSendMessage = () => {
            if (inputValue.trim() === '') {
                return;
            }

            const newMessage = {
                id: Date.now(),
                text: inputValue,
            };

            setMessages((prevMessages) => [...prevMessages, newMessage]);
            setInputValue('');
        };

    return (
        <Paper sx={{ maxWidth: 936, margin: 'auto', overflow: 'hidden' , height:'100vh', flexDirection: 'column',}}>
            <AppBar
                position="static"
                color="default"
                elevation={0}
                sx={{ borderBottom: '1px solid rgba(0, 0, 0, 0.12)' }}
            >
            </AppBar>
            <Typography sx={{ my: 5, mx: 2 }} color="text.secondary" align="center">
                欢迎来到聊天室😘😊😁😍😎
            </Typography>
            <div style={{display:'flex', flexDirection: 'column', height: '100%'}}>
                <div style={{flex: '1', maxHeight: '200px', overflowY: 'scroll' }}>
                    {messages.map((message) => (
                        <div key={message.id}>{message.text}</div>
                    ))}
                </div>
                <div style={{display: 'flex', justifyContent: 'flex-end', alignItems: 'center'}}>
                    <input type="text" value={inputValue} onChange={handleInputChange} />
                    <Button variant="contained" color="primary" onClick={handleSendMessage}>
                        Send
                    </Button>
                </div>
            </div>
        </Paper>
    );
}