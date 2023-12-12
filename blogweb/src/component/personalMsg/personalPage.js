import * as React from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/material/styles';
import {useContext, useEffect} from "react";
import UserContext from "../../identity";
import {Table, TableBody, TableCell, TableContainer, TableHead, TableRow} from "@mui/material";


const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));
const messageMap = {};

//格式化时间
const FormateData = (date) => {
    const datetime = new Date(date);

    const year = datetime.getFullYear();
    const month = String(datetime.getMonth() + 1).padStart(2, '0');
    const day = String(datetime.getDate()).padStart(2, '0');
    const hours = String(datetime.getHours()).padStart(2, '0');
    const minutes = String(datetime.getMinutes()).padStart(2, '0');
    const seconds = String(datetime.getSeconds()).padStart(2, '0');

    const formattedDatetime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedDatetime;
}
export default function PersonalPage() {
    useEffect(() => {
        // 在组件渲染完成后执行含税函数
        HandleFind();
    }, []);
    const {identity} = useContext(UserContext)
    const formData = new FormData();
    formData.append('Identity',identity);
    const HandleFind = async () => {
        try{
            const response = await fetch('http://localhost:8080/user/getUserByIdentity', {
                method: 'POST',
                data: formData,
            });
            if (response.ok) {
                const jsonData = await response.json();
                messageMap['Name'] = jsonData.message['Name'];
                messageMap['Phone'] = jsonData.message['Phone'];
                messageMap['Email'] = jsonData.message['Email'];
                const createTime = jsonData.message['CreatedAt']
                messageMap['CreatedAt'] = FormateData(createTime)
                const LastLoginTime = jsonData.message['UpdatedAt']
                messageMap['LastLoginTime'] = FormateData(LastLoginTime)
                console.log(messageMap)
            }else {
                alert("Error");
            }
        }catch (error) {
            console.log(error);
        }
    }
    return (
            <TableContainer component={Paper}>
                <Table aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>姓名</TableCell>
                            <TableCell align="right">电话</TableCell>
                            <TableCell align="right">邮箱</TableCell>
                            <TableCell align="right">创建时间</TableCell>
                            <TableCell align="right">最后登录时间</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                            <TableRow>
                                <TableCell component="th" scope="row">{messageMap.Name}</TableCell>
                                <TableCell align="right">{messageMap.Phone}</TableCell>
                                <TableCell align="right">{messageMap.Email}</TableCell>
                                <TableCell align="right">{messageMap.CreatedAt}</TableCell>
                                <TableCell align="right">{messageMap.LastLoginTime}</TableCell>
                            </TableRow>
                    </TableBody>
                </Table>
            </TableContainer>
    );
}
