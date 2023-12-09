import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Typography from '@mui/material/Typography';
import Paper from '@mui/material/Paper';


export default function GroupPage() {

    return (
        <Paper sx={{ maxWidth: 936, margin: 'auto', overflow: 'hidden' }}>
            <AppBar
                position="static"
                color="default"
                elevation={0}
                sx={{ borderBottom: '1px solid rgba(0, 0, 0, 0.12)' }}
            >
            </AppBar>
            <Typography sx={{ my: 5, mx: 2 }} color="text.secondary" align="center">
                群聊消息页面
            </Typography>
            <div>
                <div style={{ maxHeight: '200px', overflowY: 'scroll' }}>
                    <h1>welcome!</h1>
                </div>
            </div>
        </Paper>
    );
}