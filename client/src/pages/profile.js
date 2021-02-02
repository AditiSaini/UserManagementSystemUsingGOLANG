import React from 'react';
import { Container, Row, Col, Image, Card, Button, ButtonGroup } from 'react-bootstrap'
import { withRouter } from "react-router-dom";

class ProfilePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = { username: undefined, fullname: undefined, token: undefined, pic: "https://picsum.photos/id/237/200/300" }
        this.handleLogout = this.handleLogout.bind(this);
    }

    componentDidMount() {
        const token = localStorage.token;
        if (token) {
            //1. Get the profile data of the user
            fetch("http://localhost:4000/profile", {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json',
                    // Accept: 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            })
                .then(resp => resp.json())
                .then(data => {
                    this.setState({ username: data.Nickname, fullname: data.Username, token: token })
                })

            //2. Get the user profile image
            fetch("http://localhost:4000/showProfilePicture", {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json',
                    // Accept: 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            })
                .then(resp => resp.blob()).then(data => { this.setState({ pic: URL.createObjectURL(data) }) })
        }
    }

    handleLogout() {
        const token = localStorage.token;
        if (!token) {
            return
        }
        //1. Remove session token from local storage
        localStorage.clear();
        //2. Call the api to indicate logout
        fetch("http://localhost:4000/logout", {
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            }
        })
            .then(resp => resp.json())
            .then(data => {
                this.props.history.push("/")
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    render() {
        const { username, fullname, pic } = this.state;
        return (<>
            <Container style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                <Row>
                    <Card style={{ marginTop: '50px', height: '400px' }}>
                        <Card.Img variant="top" src={pic} />
                        <Card.Body>
                            <center>
                                <Card.Title >My Profile</Card.Title>
                                <Card.Text>
                                    Name: {fullname}<br />
                                    Username: {username}
                                </Card.Text>
                                <ButtonGroup className="mr-2" aria-label="First group">
                                    <Button variant="secondary" style={{ marginTop: '50px' }} onClick={() => this.props.history.push("/picture")}>Change Profile Picture</Button>
                                    <Button variant="info" style={{ marginTop: '50px' }} onClick={this.handleLogout}>Logout</Button>
                                    <Button variant="warning" style={{ marginTop: '50px' }} onClick={() => this.props.history.push("/info")}>Change Profile Information</Button>
                                </ButtonGroup>
                            </center>
                        </Card.Body>
                    </Card>
                </Row>
            </Container>
        </>);
    }
}

export default withRouter(ProfilePage);