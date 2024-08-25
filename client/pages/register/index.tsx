import { useState, useContext, useEffect } from 'react';
import { useRouter } from 'next/router';
import { AuthContext, UserInfo } from '../../modules/auth_provider';
import { API_URL } from '../../constants';

const Register = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const { authenticated } = useContext(AuthContext);

    const router = useRouter();

    useEffect(() => {
        if (authenticated) {
            router.push('/');
        }
    }, [authenticated]);

    const submitHandler = async (e: { preventDefault: () => void; }) => {
        e.preventDefault();

        try {
            const res = await fetch(`${API_URL}/auth/register`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    username, email, password
                }),
            });

            const data = await res.json();
            if (res.ok) {
                const user: UserInfo = {
                    username: data.user.username,
                    id: data.user.id,
                }
                localStorage.setItem('user_info', JSON.stringify(user))
                var userInfo = localStorage.getItem('user_info')
                console.log('User info:1', userInfo) // Debug: Check the user info
                return router.push('/')
            } else {
                // Handle registration errors
                console.error('Registration failed:', data.error);
            }
        } catch (err) {
            console.error('Error registering user:', err);
        }
    };

    return (
        <div className="flex items-center justify-center min-w-full min-h-screen">
            <form className="flex flex-col md:w-1/5">
                <div className="text-3xl font-bold text-center">
                    <span className="text-blue">Register</span>
                </div>

                <input
                    type="text"
                    placeholder="Username"
                    className="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />

                <input
                    type="email"
                    placeholder="Email"
                    className="p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />

                <input
                    type="password"
                    placeholder="Password"
                    className="p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />

                <button
                    className="p-3 mt-6 rounded-md bg-blue font-bold text-white"
                    type="submit"
                    onClick={submitHandler}
                >
                    Register
                </button>
                <p className="text-gray-500 text-center mt-4">
                    Already have an account? <a href="/login" className="text-blue underline">Login</a>
                </p>
            </form>
        </div>
    );
};

export default Register;