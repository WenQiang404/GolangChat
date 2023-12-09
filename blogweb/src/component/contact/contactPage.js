import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Typography from '@mui/material/Typography';
import Paper from '@mui/material/Paper';


export default function ContactPage() {

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
                欢迎来到的来联盟
            </Typography>
            <div>
                <div style={{ maxHeight: '200px', overflowY: 'scroll' }}>
                    <h1>welcome!</h1>
                </div>
            </div>
        </Paper>
    );
}