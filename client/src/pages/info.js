import React from 'react';
import { Container, Row, Col, Image, Card, Form, Button, ButtonGroup } from 'react-bootstrap'
import { withRouter } from "react-router-dom";

class InfoPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = { username: undefined, password: undefined, token: undefined }
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
        const token = localStorage.token;

        // stop here if form is invalid
        if (!token) {
            return;
        }

        if (username != undefined) {
            const url = "http://localhost:4000/profile/update";
            const requestOptions = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
                body: JSON.stringify({ "name": username })
            };
            fetch(url, requestOptions)
                .then(data => data.json())
                .catch((error) => {
                    console.error('Error:', error);
                });
        }
        if (password != undefined) {
            const url = "http://localhost:4000/changePassword";
            const requestOptions = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
                body: JSON.stringify({ "password": password })
            };
            fetch(url, requestOptions)
                .then(data => data.json())
                .catch((error) => {
                    console.error('Error:', error);
                });
        }
        this.props.history.push("/profile")
    }

    render() {
        return (<>
            <Container style={{ marginTop: '50px' }}>
                <Form>
                    <Form.Group controlId="formBasicEmail">
                        <Form.Label>Change your name</Form.Label>
                        <Form.Control type="text" onChange={this.handleChangeUsername} />
                    </Form.Group>
                    <Form.Group controlId="formBasicPassword">
                        <Form.Label>Change your password</Form.Label>
                        <Form.Control type="password" onChange={this.handleChangePassword} />
                    </Form.Group>
                    <Button variant="primary" type="submit" onClick={this.handleSubmit}>
                        Update
                    </Button>
                </Form></Container>
        </>);
    }
}

export default withRouter(InfoPage);