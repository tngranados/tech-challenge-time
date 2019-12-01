import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import {
    AppBar,
    Link,
    Toolbar,
    Grid
} from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        header: {
            background: 'white',
            boxShadow: 'none'
        },
        toolbar: {
            [theme.breakpoints.down('xs')]: {
                padding: theme.spacing(2, 1)
            },
            [theme.breakpoints.up('sm')]: {
                padding: theme.spacing(2, 3)
            },
            [theme.breakpoints.up('md')]: {
                padding: theme.spacing(2, 10)
            },
            [theme.breakpoints.up('lg')]: {
                padding: theme.spacing(2, 30)
            },
            [theme.breakpoints.up('xl')]: {
                padding: theme.spacing(2, 55)
            },
            justifyContent: 'space-between'
        },
        title: {
            ...theme.typography.button,
            marginRight: '2em',
            fontWeight: 600,
            transition: '0.2s',
            color: theme.palette.text.primary,
            '&:hover': {
                color: theme.palette.primary.main,
                textDecoration: 'none',
            }
        }
    })
);

export const Header: React.FC = props => {
    const classes = useStyles(props);

    return (
        <AppBar position="static" color="default" className={classes.header}>
            <Toolbar variant="dense" className={classes.toolbar}>
                <Grid
                    container
                    direction="row"
                    justify="flex-start"
                    alignItems="center"
                >
                    <Grid item xs={11} sm={9}>
                        <Grid container direction="row" justify="flex-start">
                            <Link
                                component="span"
                                className={classes.title}
                            >
                                Time Tracking
                            </Link>
                        </Grid>
                    </Grid>
                </Grid>
            </Toolbar>
        </AppBar>
    );
};
