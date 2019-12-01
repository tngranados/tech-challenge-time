import React, { useState } from 'react';
import {
  makeStyles,
  createStyles,
  Theme,
  Paper,
  Snackbar,
  IconButton,
  SnackbarContent
} from '@material-ui/core';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import { CloseRounded } from '@material-ui/icons';
import { Header } from './components/Header';
import { Sessions } from './components/Sessions';
import { CssBaseline } from '@material-ui/core';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#10C6A3',
      contrastText: '#FFFFFF'
    },
    text: {
      primary: '#333333',
    },
    background: {
      default: '#F1F1F1'
    }
  },
});

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      backgroundColor: theme.palette.background.default,
      [theme.breakpoints.down('xs')]: {
        margin: theme.spacing(2, 1)
      },
      [theme.breakpoints.up('sm')]: {
        margin: theme.spacing(2, 3)
      },
      [theme.breakpoints.up('md')]: {
        margin: theme.spacing(2, 10)
      },
      [theme.breakpoints.up('lg')]: {
        margin: theme.spacing(2, 30)
      },
      [theme.breakpoints.up('xl')]: {
        margin: theme.spacing(2, 55)
      }
    },
    page: {
      padding: theme.spacing(2, 3),
      maxWidth: '100%',
      minHeight: '85vh',
      overflowX: 'hidden'
    },
    close: {
      padding: theme.spacing(0.5)
    },
    errorSnackbar: {
      backgroundColor: theme.palette.error.main
    }
  })
);

const App: React.FC = props => {
  const classes = useStyles(props);
  const [error, setError] = useState('');

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Header />
      <div className={classes.container}>
        <Paper className={classes.page}>
          <Sessions setError={setError} />
        </Paper>
        <Snackbar
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'left'
          }}
          open={error !== ''}
          autoHideDuration={3000}
          onClose={() => setError('')}
        >
          <SnackbarContent
            className={classes.errorSnackbar}
            aria-describedby="message-id"
            message={<span id="error">{error}</span>}
            action={[
              <IconButton
                key="close"
                aria-label="close"
                color="inherit"
                className={classes.close}
                onClick={() => setError('')}
              >
                <CloseRounded />
              </IconButton>
            ]}
          />
        </Snackbar>
      </div>
    </ThemeProvider>
  );
}

export default App;
