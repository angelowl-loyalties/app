import axios from "axios";
import NextAuth from "next-auth/next";
import CredentialsProvider from "next-auth/providers/credentials";

const authOptions = {
    providers: [
        CredentialsProvider({
            name: "credentials",
            credentials: {
                email: { label: "email", type: "email", placeholder: "john@doe.com.sg" },
                password: { label: "password", type: "password" }
            },
            authorize: async (credentials) => {
                // Make POST request to log 
                const request = axios.post(process.env.NEXT_PUBLIC_API_URL + "AUTH_ENDPOINT_HERE", {
                    email: credentials.email,
                    password: credentials.password
                })

                // Store token response
                return await request.then((response) => {
                    const user = { id: "Token " + response.data.token, token: "Token " + response.data.token }
                    return user
                }).catch((e) => {
                    console.log(e);
                    return null;
                });
            },
        }),
    ],
    callbacks: {
        jwt: async ({ token, user }) => {
            if (user) {
                token.id = user.id;
            }

            return token;
        },
        session: ({ session, token }) => {
            if (token) {
                session.id = token.id;
                session.userId = token.id;
            }

            return session;
        },
    },
    secret: "ABC123",
    jwt: {
        secret: "ABC123",
    },
    pages: {
        signIn: "/login",
        signOut: "/login",
    },
};

export default NextAuth(authOptions);
