import React from 'react';
import { Container, Row, Col, Image, Card, Form, Button, ButtonGroup } from 'react-bootstrap'
import { withRouter } from "react-router-dom";

class PicturePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = { username: undefined, password: undefined, token: undefined, file: null, rawFile: null }
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleChange(event) {
        this.setState({
            file: URL.createObjectURL(event.target.files[0]),
            rawFile: event.target.files[0]
        })
    }

    handleSubmit() {
        const token = localStorage.token;
        // stop here if form is invalid
        if (!token) {
            return;
        }

        let formData = new FormData();
        formData.append(
            'myFile', this.state.rawFile);

        fetch('http://localhost:4000/uploadProfilePicture', {
            method: 'POST', headers: { 'Authorization': `Bearer ${token}` }
            , body: formData
        })
            .then(response => { response.json(); this.props.history.push("/profile") })
            .catch((error) => {
                console.error('Error:', error);
            });
        ;
    }

    render() {
        return (<>
            <Container style={{ marginTop: '50px' }}>
                <h2>Upload your new profile picture</h2>
                <input type="file" onChange={this.handleChange} />
                <img src={this.state.file} />
                <Button variant="info" style={{ marginTop: '50px' }} onClick={this.handleSubmit}>Update Picture</Button>
            </Container>
        </>);
    }
}

export default withRouter(PicturePage);