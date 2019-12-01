import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { SessionsApi } from '../api/Sessions';
import { Session as SessionModel } from '../models/Session';
import { Grid, Tooltip } from '@material-ui/core';
import Icon from '@material-ui/core/Icon';
import { PauseRounded, ClearRounded } from '@material-ui/icons';
import Moment from 'react-moment';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        row: {
            minHeight: '2rem',
        },
        button: {
            cursor: 'pointer',
            transition: '0.2s',
            color: theme.palette.text.primary,
            '&:hover': {
                color: theme.palette.primary.main
            }
        }
    })
);

const api = new SessionsApi();

interface Props {
    session: SessionModel;
    setSession: (session: SessionModel) => void;
    deleteSession: (session: SessionModel) => void;
    setError: (error: string) => void;
}

export const Session: React.FC<Props> = props => {
    const classes = useStyles(props);
    const { session, setSession, deleteSession, setError } = props;

    const startedAt = new Date(session.StartedAt);
    const finishedAt = new Date(session.FinishedAt);
    const finished = finishedAt > new Date(0);
    const finishedSameDay = startedAt.getFullYear() === finishedAt.getFullYear() && startedAt.getMonth() === finishedAt.getMonth() && startedAt.getDay() === finishedAt.getDay();

    const handleDelete = (event: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
        api.delete(session.ID)
            .then(() => {
                deleteSession(session);
            })
            .catch(e => setError(`Error deleting session: ${e}`))
    }

    const handlePause = (event: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
        // Copy session
        let newSession: SessionModel = JSON.parse(JSON.stringify(session));
        newSession.FinishedAt = new Date();

        api.update(newSession)
            .then(() => {
                setSession(newSession);
            })
            .catch(e => setError(`Error updating session: ${e}`))
    }

    return (
        <Grid container direction='row' className={classes.row} justify='space-between'>
            <Grid item xs={2}>
                {session.Name}
            </Grid>
            <Grid item xs={6}>
                {startedAt.toLocaleString()}
                {' - '}
                {finished ? finishedSameDay ? finishedAt.toLocaleTimeString() : finishedAt.toLocaleString() : 'Running'}
            </Grid>
            <Grid item>
                {finished ? (
                    <Moment interval={0} duration={startedAt} date={finishedAt} />
                ) : (
                        <Moment interval={1000} duration={startedAt} durationFromNow />
                    )}
            </Grid>
            <Grid item>
                <Tooltip title={finished ? 'Delete' : 'Stop'} placement='right'>
                    <Icon className={classes.button} onClick={finished ? handleDelete : handlePause}>
                        {finished ? <ClearRounded /> : <PauseRounded />}
                    </Icon>
                </Tooltip>
            </Grid>
        </Grid>
    );
};
