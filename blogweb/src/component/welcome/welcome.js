import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import {
    AppBar,
    Toolbar,
    Typography,
    Button,
    Container,
    Grid,
    Card,
    CardActionArea,
    CardMedia,
    CardContent,
    CardActions,
    Link,
    Box,
    IconButton,
    BottomNavigation,
    BottomNavigationAction
} from '@material-ui/core';
import {
    Favorite as FavoriteIcon,
    Share as ShareIcon,
    MoreVert as MoreVertIcon
} from '@material-ui/icons';
import { useHistory } from 'react-router-dom';

const useStyles = makeStyles(theme => ({
    root: {
        flexGrow: 1,
        backgroundColor: theme.palette.background.paper,
        minHeight: '100vh'
    },
    appBar: {
        backgroundColor: theme.palette.primary.main
    },
    title: {
        flexGrow: 1
    },
    heroContent: {
        padding: theme.spacing(8, 0, 6),
        background: `linear-gradient(to bottom, ${theme.palette.primary.light}, ${theme.palette.primary.dark})`
    },
    heroButtons: {
        marginTop: theme.spacing(4)
    },
    card: {
        maxWidth: 345,
        margin: theme.spacing(2)
    },
    media: {
        height: 140
    },
    footer: {
        backgroundColor: theme.palette.background.paper,
        padding: theme.spacing(6)
    },
    button: {
        marginTop: theme.spacing(2),
        backgroundColor: theme.palette.secondary.main,
        color: theme.palette.secondary.contrastText
    }
}));

const WelcomePage = () => {
    const classes = useStyles();
    const history = useHistory();

    const handleButtonClick = () => {
        history.push('/hello');
    };

    return (
        <div className={classes.root}>
            <AppBar position="static" className={classes.appBar}>
                <Toolbar>
                    <Typography variant="h6" className={classes.title}>
                        Welcome
                    </Typography>
                    <Button color="inherit">Login</Button>
                </Toolbar>
            </AppBar>
            <div className={classes.heroContent}>
                <Container maxWidth="sm">
                    <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                        Welcome to our website
                    </Typography>
                    <Typography variant="h5" align="center" color="textSecondary" paragraph>
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vel augue euismod, imperdiet ex vel, pharetra
                        quam. Aliquam vitae est molestie, vulputate augue vel, dapibus enim. Praesent pharetra elit ut elit posuere,
                        sed viverra ante blandit.
                    </Typography>
                    <div className={classes.heroButtons}>
                        <Grid container spacing={2} justify="center">
                            <Grid item>
                                <Button variant="contained" className={classes.button} onClick={handleButtonClick}>
                                    Go to Hello page
                                </Button>
                            </Grid>
                        </Grid>
                    </div>
                </Container>
            </div>
            <Container maxWidth="md">
                <Grid container spacing={4}>
                    <Grid item xs={12} sm={6} md={4}>
                        <Card className={classes.card}>
                            <CardActionArea>
                                <CardMedia
                                    className={classes.media}
                                    image="https://source.unsplash.com/random"
                                    title="Image title"
                                />
                                <CardContent>
                                    <Typography gutterBottom variant="h5" component="h2">
                                        Card 1
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary" component="p">
                                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vel augue euismod, imperdiet ex vel,
                                        pharetra quam.
                                    </Typography>
                                </CardContent>
                            </CardActionArea>
                            <CardActions>
                                <IconButton>
                                    <FavoriteIcon />
                                </IconButton>
                                <IconButton>
                                    <ShareIcon />
                                </IconButton>
                                <IconButton>
                                    <MoreVertIcon />
                                </IconButton>
                            </CardActions>
                        </Card>
                    </Grid>
                    <Grid item xs={12} sm={6} md={4}>
                        <Card className={classes.card}>
                            <CardActionArea>
                                <CardMedia
                                    className={classes.media}
                                    image="https://source.unsplash.com/random"
                                    title="Image title"
                                />
                                <CardContent>
                                    <Typography gutterBottom variant="h5" component="h2">
                                        Card 2
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary" component="p">
                                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vel augue euismod, imperdiet ex vel,
                                        pharetra quam.
                                    </Typography>
                                </CardContent>
                            </CardActionArea>
                            <CardActions>
                                <IconButton>
                                    <FavoriteIcon />
                                </IconButton>
                                <IconButton>
                                    <ShareIcon />
                                </IconButton>
                                <IconButton>
                                    <MoreVertIcon />
                                </IconButton>
                            </CardActions>
                        </Card>
                    </Grid>
                    <Grid item xs={12} sm={6} md={4}>
                        <Card className={classes.card}>
                            <CardActionArea>
                                <CardMedia
                                    className={classes.media}
                                    image="https://source.unsplash.com/random"
                                    title="Image title"
                                />
                                <CardContent>
                                    <Typography gutterBottom variant="h5" component="h2">
                                        Card 3
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary" component="p">
                                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vel augue euismod, imperdiet ex vel,
                                        pharetra quam.
                                    </Typography>
                                </CardContent>
                            </CardActionArea>
                            <CardActions>
                                <IconButton>
                                    <FavoriteIcon />
                                </IconButton>
                                <IconButton>
                                    <ShareIcon />
                                </IconButton>
                                <IconButton>
                                    <MoreVertIcon />
                                </IconButton>
                            </CardActions>
                        </Card>
                    </Grid>
                </Grid>
            </Container>
            <Box mt={8}>
                <BottomNavigation
                    value={value}
                    onChange={(event, newValue) => {
                        setValue(newValue);
                    }}
                    showLabels
                >
                    <BottomNavigationAction label="About" />
                    <BottomNavigationAction label="Contact" />
                    <BottomNavigationAction label="Privacy Policy" />
                </BottomNavigation>
            </Box>
        </div>
    );
};

export default WelcomePage;
