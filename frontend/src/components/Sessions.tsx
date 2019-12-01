import React, { useEffect, useState, useCallback } from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { SessionsApi } from '../api/Sessions';
import { Session as SessionModel } from '../models/Session';
import { Session } from './Session';
import { Grid, TextField, Tooltip, Chip } from '@material-ui/core';
import Icon from '@material-ui/core/Icon';
import { PlayArrowRounded } from '@material-ui/icons';
import clsx from 'clsx';
import moment from 'moment';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        topBar: {
            marginBottom: theme.spacing(2)
        },
        button: {
            cursor: 'pointer',
            transition: '0.2s',
            color: theme.palette.text.primary,
            '&:hover': {
                color: theme.palette.primary.main
            }
        },
        filter: {
            cursor: 'pointer',
            transition: '0.2s',
            '&:hover': {
                color: 'white',
                backgroundColor: theme.palette.text.primary
            }
        },
        selectedFilter: {
            color: 'white',
            backgroundColor: `${theme.palette.primary.main} !important`
        }
    })
);

const api = new SessionsApi();

interface Props {
    setError: (error: string) => void;
}

export const Sessions: React.FC<Props> = props => {
    const classes = useStyles(props);
    const { setError } = props;
    const [sessions, setSessions] = useState<SessionModel[]>();
    const [filteredSessions, setFilteredSessions] = useState<SessionModel[]>();
    const [newSessionName, setNewSessionName] = useState('');
    const [filter, setFilter] = useState('');

    const setSessionFunc = (id: number) => {
        return (s: SessionModel) => {
            if (sessions) {
                // Make a copy of sessions
                const newSessions: SessionModel[] = JSON.parse(JSON.stringify(sessions));

                const index = newSessions.findIndex(el => el.ID === s.ID);
                newSessions[index] = s;
                setSessions(newSessions);
            }
        }
    }

    const deleteSessionFunc = (id: number) => {
        return (s: SessionModel) => {
            if (sessions) {
                // Make a copy of sessions
                const newSessions: SessionModel[] = JSON.parse(JSON.stringify(sessions));

                const index = newSessions.findIndex(el => el.ID === s.ID);
                newSessions.splice(index, 1);
                setSessions(newSessions);
            }
        }
    }

    const handleOnChangeNewSession = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewSessionName(event.currentTarget.value);
    }

    const handleEnter = (event: React.KeyboardEvent) => {
        if (event.key === 'Enter') {
            handleAddSession();
        }
    }

    const handleAddSession = () => {
        if (newSessionName !== '') {
            const newSession: SessionModel = {
                ID: 0,
                Name: newSessionName,
                StartedAt: new Date(),
                FinishedAt: new Date(0),
            }
            api.add(newSession)
                .then(() => {
                    setNewSessionName('');
                    updateSessions();
                })
                .catch(e => setError(e));
        }
    }

    const updateSessions = useCallback(() => {
        api.getAll()
            .then(res => setSessions(res))
            .catch(e => setError(e));
    }, [setError])

    useEffect(() => {
        updateSessions();
    }, [updateSessions]);

    const handleFilterClick = (f: string) => {
        return () => {
            if (filter === f) {
                setFilter('');
            } else {
                setFilter(f);
            }
        }
    }

    useEffect(() => {
        if (sessions) {
            // Copy sessions
            const copiedSessions: SessionModel[] = JSON.parse(JSON.stringify(sessions));

            let filtered: SessionModel[];
            const now = moment(new Date())
            switch (filter) {
                case 'day': {
                    filtered = copiedSessions.filter(s => {
                        const started = moment(s.StartedAt)
                        return started.year() === now.year() && started.dayOfYear() === now.dayOfYear()
                    });
                    break;
                }
                case 'week': {
                    filtered = copiedSessions.filter(s => {
                        const started = moment(s.StartedAt)
                        return started.year() === now.year() && started.week() === now.week()
                    });
                    break;
                }
                case 'month': {
                    filtered = copiedSessions.filter(s => {
                        const started = moment(s.StartedAt)
                        return started.year() === now.year() && started.month() === now.month()
                    });
                    break;
                }
                default: {
                    filtered = copiedSessions;
                    break;
                }
            }
            setFilteredSessions(filtered);
        }
    }, [filter, sessions])

    return (
        <Grid container direction='column'>
            <Grid item className={classes.topBar}>
                <Grid container direction='row' justify='space-between' alignItems='center'>
                    <Grid item xs={4}>
                        <Grid container spacing={1} alignItems='flex-end'>
                            <Grid item>
                                <TextField label='Name' value={newSessionName} onChange={handleOnChangeNewSession} onKeyUp={handleEnter} />
                            </Grid>
                            <Grid item>
                                <Tooltip title='Add' placement='right'>
                                    <Icon className={classes.button} onClick={handleAddSession}>
                                        <PlayArrowRounded />
                                    </Icon>
                                </Tooltip>
                            </Grid>
                        </Grid>
                    </Grid>
                    <Grid item xs={4}>
                        <Grid container direction='row' justify='flex-end' spacing={1}>
                            <Grid item>
                                <Chip onClick={handleFilterClick('day')} label='Day' size="small" className={filter === 'day' ? clsx(classes.filter, classes.selectedFilter) : classes.filter} />
                            </Grid>
                            <Grid item>
                                <Chip onClick={handleFilterClick('week')} label='Week' size="small" className={filter === 'week' ? clsx(classes.filter, classes.selectedFilter) : classes.filter} />
                            </Grid>
                            <Grid item>
                                <Chip onClick={handleFilterClick('month')} label='Month' size="small" className={filter === 'month' ? clsx(classes.filter, classes.selectedFilter) : classes.filter} />
                            </Grid>
                        </Grid>
                    </Grid>
                </Grid>
            </Grid>
            <Grid item>
                {filteredSessions && filteredSessions.map(s => {
                    return (
                        <Session key={`session_${s.ID}`} session={s} setSession={setSessionFunc(s.ID)} deleteSession={deleteSessionFunc(s.ID)} setError={setError} />
                    )
                })}
            </Grid>
        </Grid>
    );
};
