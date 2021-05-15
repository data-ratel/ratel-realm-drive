import React from 'react';
import { Form, Button } from 'react-bootstrap';
import './Login.css';

function Login() {
    return(
        <div className='Login-Layout'>
            <div className='Login-Card'>
                <div className='Login-Card-Title'>
                    <p>Sign in to Ratel Drive</p>
                </div>
                
                <div className='Login-Card-Body'>
                    <Form>
                        <Form.Group controlId="formCredential" className='Login-Card-Body-Cred'>
                            <Form.Control type="email" size='lg' placeholder="Email" />
                            <Form.Control type="password" size='lg' placeholder="Password" />
                        </Form.Group>

                        <Form.Group controlId="formUtils" className='Login-Card-Body-Utils'>
                            <Form.Check type="checkbox" label="Save password" />
                            
                            <Form.Text className="text-muted">
                            <a href='/'>Forgot your password?</a>
                            </Form.Text>
                        </Form.Group>

                        <Form.Group controlId="formUtils" className='Login-Card-Body-Submit'>
                            <Button variant="primary" type="submit" size='lg' className='Login-Card-Body-BtnSubmit'>
                                Submit
                            </Button>
                        </Form.Group>
                    </Form>
                </div>
            </div>
        </div>
    )
}

export default Login;