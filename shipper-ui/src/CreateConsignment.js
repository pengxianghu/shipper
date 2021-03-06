import React from 'react';
import _ from 'lodash';

class CreateConsignment extends React.Component {

    constructor(props) {
        super(props);
        this.state = {};
    }

    state = {
        created: false,
        description: '',
        weight: 0,
        containers: [],
        consignments: [],
    }

    componentWillMount() {
        const token = sessionStorage.getItem('token');
        fetch(`http://www.pengxianghu.com/rpc`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token,
            },
            body: JSON.stringify({
                service: 'go.micro.srv.consignment',
                method: 'ConsignmentService.GetConsignments',
                request: {},
            })
        })
            .then(req => req.json())
            .then((res) => {
                console.log("create consignment component mount: " + res);
                this.setState({
                    consignments: res.consignments,
                });
            });
    }

    create = () => {
        const consignment = this.state;
        const token = sessionStorage.getItem('token');
        fetch(`http://www.pengxianghu.com/rpc`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token,
            },
            body: JSON.stringify({
                service: 'go.micro.srv.consignment',
                method: 'ConsignmentService.CreateConsignment',
                request: _.omit(consignment, 'created', 'consignments'),
            }),
        })
            .then((res) => res.json())
            .then((res) => {
                if (this.state.consignments === undefined ) {
                    this.setState({
                        created: res.created,
                        consignments: [consignment],
                    });
                } else {
                    this.setState({
                        created: res.created,
                        consignments: [...this.state.consignments, consignment],
                    });
                }
            });
    }

    addContainer = e => {
        this.setState({
            containers: [...this.state.containers, e.target.value],
        });
    }

    setDescription = e => {
        this.setState({
            description: e.target.value,
        });
    }

    setWeight = e => {
        this.setState({
            weight: Number(e.target.value),
        });
    }

    render() {
        const { consignments, } = this.state;
        return (
            <div className='consignment-screen'>
                <div className='consignment-form container'>
                    <br />
                    <div className='form-group'>
                        <textarea onChange={this.setDescription} className='form-control' placeholder='Description'></textarea>
                    </div>
                    <div className='form-group'>
                        <input onChange={this.setWeight} type='number' placeholder='Weight' className='form-control' />
                    </div>
                    <br />
                    <button onClick={this.create} className='btn btn-primary'>添加</button>
                    <br />
                    <hr />
                </div>
                {(consignments && consignments.length > 0
                    ? <div className='consignment-list'>
                        <h2>Consignments</h2>
                        <hr />
                        {consignments.map((item) => (
                            <div>
                                <p>Description: {item.description}</p>
                                <p>Weight: {item.weight}</p>
                                <hr />
                            </div>
                        ))}
                    </div>
                    : false)}
            </div>
        );
    }
}

export default CreateConsignment;