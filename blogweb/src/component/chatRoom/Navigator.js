import * as React from 'react';
import Divider from '@mui/material/Divider';
import Drawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import Box from '@mui/material/Box';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import HomeIcon from '@mui/icons-material/Home';
import PeopleIcon from '@mui/icons-material/People';
import DnsRoundedIcon from '@mui/icons-material/DnsRounded';
import PermMediaOutlinedIcon from '@mui/icons-material/PhotoSizeSelectActual';
import PublicIcon from '@mui/icons-material/Public';
import personalPage from "../personalMsg/personalPage";
import groupPage from "./groupPage";
import Content from "./Content";
import ContactPage from "./contactPage";

const categories = [
    {
        id: '❤️❤️❤️',
        children: [
            {
                id: '个人资料',
                icon: <PeopleIcon />,
                active: true,
                className:'personal',
            },
            { id: '群聊', icon: <DnsRoundedIcon />,className:'group', },
            { id: '图片', icon: <PermMediaOutlinedIcon />,className:'picture', },
            { id: '联系人', icon: <PublicIcon />,className:'contact', },
        ],
    },
];

const item = {
    py: '2px',
    px: 3,
    color: 'rgba(255, 255, 255, 0.7)',
    '&:hover, &:focus': {
        bgcolor: 'rgba(255, 255, 255, 0.08)',
    },
};

const itemCategory = {
    boxShadow: '0 -1px 0 rgb(255,255,255,0.1) inset',
    py: 1.5,
    px: 3,
};

export default function Navigator(props) {
    const { ...other } = props;
    const defaultnavigator = ""
    const [selectOption, setSelectedOption] = React.useState(defaultnavigator);
    const handleClick = (option) => {
        setSelectedOption(option);
    }
    const renderPage = (selectOption) => {
        switch (selectOption) {
            // case 'personal':
            //     return <personalPage/>;
            //     break;
            // case 'picture':
            //     contentcomponent = <picturePage/>;
            //     break;
            // case 'group':
            //     contentcomponent = <groupPage/>;
            //     break;
            case 'contact':
                return <ContactPage/>;
            case 'room':
                return <Content/>;
    }


    };
    return (
        <Drawer variant="permanent" {...other}>
            <List disablePadding>
                <ListItem sx={{ ...item, ...itemCategory, fontSize: 22, color: '#fff' }}>
                    聊天室
                </ListItem>
                <ListItem sx={{ ...item, ...itemCategory }}>
                    <ListItemIcon>
                        <HomeIcon />
                    </ListItemIcon>
                    <ListItemText>聊天大厅</ListItemText>
                </ListItem>
                {categories.map(({ id, children}) => (
                    <Box key={id} sx={{ bgcolor: '#101F33' }}>
                        <ListItem sx={{ py: 2, px: 3 }}>
                            <ListItemText sx={{ color: '#fff' }}>{id}</ListItemText>
                        </ListItem>
                        {children.map(({ id: childId, icon, active ,className}) => (
                            <ListItem disablePadding key={childId}>
                                <ListItemButton selected={active} sx={item} onclick={() => handleClick(setSelectedOption(className))}>
                                    <ListItemIcon>{icon}</ListItemIcon>
                                    <ListItemText>{childId}</ListItemText>
                                </ListItemButton>
                            </ListItem>
                        ))}
                        <div>
                            {renderPage(selectOption)}
                        </div>
                        <Divider sx={{ mt: 2 }} />
                    </Box>

                ))}
            </List>
        </Drawer>
    );
}