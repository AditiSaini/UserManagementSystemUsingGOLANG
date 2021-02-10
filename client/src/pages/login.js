import React from 'react';
import './login.css';
import { withRouter } from "react-router-dom";
import { FormText } from 'react-bootstrap';



class LoginPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = { username: undefined, password: undefined, invalid: false };

        this.handleChangeUsername = this.handleChangeUsername.bind(this);
        this.handleChangePassword = this.handleChangePassword.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChangeUsername(event) {
        this.setState({ username: event.target.value });
    }

    handleChangePassword(event) {
        this.setState({ password: event.target.value });
    }

    async handleSubmit(e) {
        e.preventDefault();
        const { username, password } = this.state;

        // stop here if form is invalid
        if (!(username && password)) {
            this.setState({ invalid: true })
            return;
        }

        const url = "http://localhost:4000/login";
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json', },
            body: JSON.stringify({ "username": username, "password": password })
        };
        fetch(url, requestOptions)
            .then(data => data.json())
            .then(text => {
                if (text.access_token == "Invalid Credentials") {
                    this.setState({ invalid: true })
                } else {
                    localStorage.clear();
                    localStorage.setItem('token', text.access_token); this.props.history.push("/profile");
                }
            })
            .catch((error) => {
                this.setState({ invalid: true })
                console.error('Error:', error);
            });
    }

    render() {
        const { invalid } = this.state;
        return (<>
            <div className="login-wrapper">
                <h1>Login</h1>
                <form onSubmit={this.handleSubmit}>
                    <center>
                        <label>
                            <p>Username</p>
                            <input type="text" value={this.state.username} onChange={this.handleChangeUsername} />
                        </label>
                        <label>
                            <p>Password</p>
                            <input type="password" value={this.state.password} onChange={this.handleChangePassword} />
                        </label>
                        <div>
                            <button type="submit">Submit</button>
                        </div>
                        {invalid && (<div style={{ marginTop: '20px', fontSize: '10px' }}><i>The username and passwrd does not exist!</i></div>)}
                    </center>
                </form>
            </div>
        </>);
    }
}

export default withRouter(LoginPage);