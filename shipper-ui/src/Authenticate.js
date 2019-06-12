import React from 'react';
// import axios from 'axios';

class Authenticate extends React.Component {

    constructor(props) {
        super(props);
        this.state = {};
    }

    state = {
        authenticated: false,
        email: '',
        password: '',
        err: '',
    }

    login = () => {

        fetch(`http://www.pengxianghu.com/rpc`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                service: 'go.micro.srv.user',
                method: 'UserService.Auth',
                request: {
                    email: this.state.email,
                    password: this.state.password,
                },
            }),
        })
            .then(res => res.json())
            .then(data => {
                // console.log(data);
                this.setState({
                    token: data.token,
                    authenticated: true,
                });
                sessionStorage.setItem('token', data.token);
                this.props.onAuth(data.token);
            })
            .catch(err => this.setState({ err, authenticated: false, }));

        // let res = '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiNjY0ODNmZjYtOWVhMC00M2UwLThiNmYtMzQyYTJlM2I3OTUwIiwibmFtZSI6ImhlbnJ5IiwiZW1haWwiOiJoZW5yeSIsInBhc3N3b3JkIjoiJDJhJDEwJEd2YXpqSDZwelE1N3M1ZHZhb2ZvTU9LMGsuV3lQNjhsdTJYaFRsWnhRN3ZqMDhrMmUzOWlTIn0sImV4cCI6MTU2MDQxNDA3MiwiaXNzIjoiZ28ubWljcm8uc3J2LnVzZXIifQ.eqZKQm0oi2XA8OpHhk4-_E59mekjmSVZC41gnrtAiHs"}';

        // let obj = JSON.parse(res);
        // window.localStorage.setItem('token', obj.token);
        // this.props.onAuth();

    }

    signup = () => {
        fetch(`http://www.pengxianghu.com/rpc`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                service: 'go.micro.srv.user',
                method: 'UserService.Create',
                request: {
                    email: this.state.email,
                    password: this.state.password,
                    name: this.state.name,
                },
            }),
        })
            .then((res) => res.json())
            .then((res) => {
                // console.log(res);
                // this.setState({
                //     authenticated: true,
                // });
                // this.props.onAuth(res.token);
            })
            .catch(err => this.setState({ err, authenticated: false, }));
    }

    setEmail = e => {
        this.setState({
            email: e.target.value,
        });
    }

    setPassword = e => {
        this.setState({
            password: e.target.value,
        });
    }

    setName = e => {
        this.setState({
            name: e.target.value,
        });
    }

    render() {
        return (
            <div className='Authenticate'>
                <div className='Login'>
                    <div className='form-group'>
                        <input
                            type="email"
                            onChange={this.setEmail}
                            placeholder='E-Mail'
                            className='form-control' />
                    </div>
                    <div className='form-group'>
                        <input
                            type="password"
                            onChange={this.setPassword}
                            placeholder='Password'
                            className='form-control' />
                    </div>
                    <button className='btn btn-primary' onClick={this.login}>登录</button>
                    <br /><br />
                </div>
                <div className='Sign-up'>
                    <div className='form-group'>
                        <input
                            type='input'
                            onChange={this.setName}
                            placeholder='Name'
                            className='form-control' />
                    </div>
                    <div className='form-group'>
                        <input
                            type='email'
                            onChange={this.setEmail}
                            placeholder='E-Mail'
                            className='form-control' />
                    </div>
                    <div className='form-group'>
                        <input
                            type='password'
                            onChange={this.setPassword}
                            placeholder='Password'
                            className='form-control' />
                    </div>
                    <button className='btn btn-primary' onClick={this.signup}>注册</button>
                </div>
            </div>
        );
    }
}

export default Authenticate;