import React, {useState} from 'react';
import { Form, Button } from 'react-bootstrap';
import './Login.css';

function Login() {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    return(
        <div className='Login-Layout'>
            <div className='Login-Card'>
                <div className='Login-Card-Title'>
                    <p>Sign in to Ratel Drive</p>
                </div>
                
                <div className='Login-Card-Body'>
                    <Form>
                        <Form.Group controlId="formCredential" className='Login-Card-Body-Cred'>
                            <Form.Control type="email" size='lg' placeholder="Email" value={email} onChange={e=> setEmail(e.target.value)}/>
                            <Form.Control type="password" size='lg' placeholder="Password" value={password} onChange={e=> setPassword(e.target.value)}/>
                        </Form.Group>

                        <Form.Group controlId="formUtils" className='Login-Card-Body-Utils'>
                            <Form.Check type="checkbox" label="Save password" />
                            
                            <Form.Text className="text-muted">
                            <a href='/'>Forgot your password?</a>
                            </Form.Text>
                        </Form.Group>

                        <Form.Group controlId="formUtils" className='Login-Card-Body-SignIn'>
                            <Button 
                                variant="primary" 
                                type="submit" 
                                size='lg' 
                                className='Login-Card-Body-BtnSignIn'
                                onClick={e => onLogin(e, {email, password})}
                            >
                                SIGN IN
                            </Button>
                        </Form.Group>
                    </Form>
                </div>
            </div>
        </div>
    )
}

export default Login;

interface LoginProps {
    email: string,
    password: string
}
async function onLogin(e: React.MouseEvent<HTMLElement, MouseEvent>, props: LoginProps) {
    e.preventDefault();

    const req_options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(
            {
                user: {
                    email: props.email,
                    password: props.password
                }
            }
        )
    };

    fetch('/api/login', req_options)
        .then(rsp => rsp.json())
        .then(data => {
            console.log(data)
        });
}