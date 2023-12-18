import {createTheme, ThemeProvider} from "@mui/material/styles";
import Grid from "@mui/material/Grid";
import CssBaseline from "@mui/material/CssBaseline";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Avatar from "@mui/material/Avatar";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import FormControlLabel from "@mui/material/FormControlLabel";
import Checkbox from "@mui/material/Checkbox";
import Button from "@mui/material/Button";
import Link from "@mui/material/Link";
import * as React from "react";
import {useState} from "react";
import MyImg from '../img/index.jpg'
import CryptoJS from 'crypto-js'
import {useNavigate} from "react-router-dom";
import identity from "../../identity";
import {alertTitleClasses} from "@mui/material";

function Copyright(props) {

    return (
        <Typography variant="body2" color="text.secondary" align="center" {...props}>
            {'Copyright © '}
            <Link color="inherit" href="https://mui.com/">
                聊天室
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    );
}

const defaultTheme = createTheme();

export default function Register() {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [repassword, setRePassword] = useState('');
    const [email, setEmail] = useState('');
    const [isValid, setIsValid] = useState(true);       //邮箱格式校验
    const [verify, setVerify] = useState('')            //邮箱验证码确认
    const [showIcon, setShowIcon] = useState(false)
    let verifyCode = ""
    const navigate = useNavigate();

    const handleEmailChange = (event) => {
        const inputEmail = event.target.value;
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        setIsValid(emailRegex.test(inputEmail));
        setEmail(inputEmail);
    };

    const checkEmailValidation = async (event) => {
        event.preventDefault();
        const formDataEmail = new FormData();
        formDataEmail.append('email',email);
        try{
            const response =  await fetch('http://localhost:8080/checkEmail', {
                method: 'POST',
                body: formDataEmail,
            });
            if (response.ok) {
                const jsonData = await response.json(); //获取响应json中的数据
                const emailStatusCode = jsonData["code"];

                if (emailStatusCode === -1) {
                    alert("邮箱格式错误")

                }else{
                    verifyCode =emailStatusCode;
                }
            } else {
                alert('邮箱校验失败');
                console.log('Error:', response.status);
            }
        }catch (error) {
            console.log("Error to fetch:", error)
        }
    }

    const HandleSubmit = async (event) => {

        const formData = new FormData();
        if (username === null) {
            alert("请输入用户名");

        }else if (password !== repassword) {
            alert("两次密码不一致！");

        }else if (verifyCode !== verify) {
            alert("验证码不正确！")
            setShowIcon(false);

        }else {
            setShowIcon(true)
            //加盐加密
            const salt = 'thisisasaltmessage';
            const encryptData = CryptoJS.SHA256(salt + password).toString();
            formData.append('name',username);
            formData.append('password',encryptData);
            formData.append('email',email);
            try{
                const response = await fetch('http://localhost:8080/createUser', {
                    method: 'POST',
                    body: formData,
                });
                if (response.ok) {
                    const jsonData = await response.json(); //获取响应json中的数据
                    const code =jsonData.data["code"];
                    if (code === -1) {
                        alert("用户名已经注册！")
                    }else if (code === 1) {

                        localStorage.setItem('token', identity);
                        // 执行页面跳转
                        navigate('/');
                    }else {
                        alert("return code is error")
                    }

                } else {
                    alert('Invalid credentials');
                    console.log('Error:', response.status);
                }
            }catch (error) {
                console.log("Error to fetch:", error)
            }
        }
    };

    return (
        <ThemeProvider theme={defaultTheme}>
            <Grid container component="main" sx={{ height: '100vh' }}>
                <CssBaseline />
                <Grid
                    item
                    xs={false}
                    sm={4}
                    md={7}
                    sx={{
                        backgroundImage: `url(${MyImg})`,
                        backgroundRepeat: 'no-repeat',
                        backgroundColor: (t) =>
                            t.palette.mode === 'light' ? t.palette.grey[50] : t.palette.grey[900],
                        backgroundSize: 'cover',
                        backgroundPosition: 'center',
                    }}
                />
                <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
                    <Box
                        sx={{
                            my: 8,
                            mx: 4,
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                        }}
                    >
                        <Avatar sx={{ m: 1, bgcolor: 'secondary.main' ,alt: "admin" ,src:"/img/1.png"}}>
                            <LockOutlinedIcon />
                        </Avatar>
                        <Typography component="h1" variant="h5">
                            Register
                        </Typography>
                        <Box component="form" noValidate  sx={{ mt: 1 }}>
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                id="Username"
                                label="Username"
                                value={username}
                                autoComplete="Username"
                                autoFocus
                                onChange={(event) => setUsername(event.target.value)}
                            />
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="password"
                                label="password"
                                type="password"
                                value={password}
                                id="password"
                                autoComplete="current-password"
                                onChange={(event) => setPassword(event.target.value)}
                            />
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="repassword"
                                label="password"
                                type="password"
                                value={repassword}
                                id="repassword"
                                autoComplete="current-repassword"
                                onChange={(event) => setRePassword(event.target.value)}
                            />
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="email"
                                label="email"
                                type="email"
                                value={email}
                                id="phone"
                                autoComplete="current-email"
                                onChange={handleEmailChange}
                            >
                                {!isValid && <p>Please enter a valid email address</p>} </TextField>
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="verify"
                                label="verify"
                                type="verify"
                                value={verify}
                                id="password"
                                autoComplete="current-verify"
                                onChange={(event) => setVerify(event.target.value)}
                            />
                            <Button
                                type="submit"
                                variant="contained"
                                sx={{ mt: 2, mb: 1 }}
                                onClick={checkEmailValidation}
                            >
                                Send Email
                                {showIcon && <img src={showIcon ? 'img/yes.png' : 'img/no.png'} alt="图标"/>}
                            </Button>
                            <br/>
                            <FormControlLabel
                                control={<Checkbox value="remember" color="primary" />}
                                label="Remember me"
                            />
                            <Button
                                type="submit"
                                fullWidth
                                variant="contained"
                                sx={{ mt: 3, mb: 2 }}
                                onSubmit={HandleSubmit}
                            >
                               Register
                            </Button>
                            <Grid container>
                                <Grid item xs>
                                    <Link href="/" variant="body2">
                                        {"Already has account? Return to login"}
                                    </Link>
                                </Grid>
                            </Grid>
                            <Copyright sx={{ mt: 5 }} />
                        </Box>
                    </Box>
                </Grid>
            </Grid>
        </ThemeProvider>
    );
}
