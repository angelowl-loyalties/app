import axios from 'axios';
import NextAuth from 'next-auth/next';
import CredentialsProvider from 'next-auth/providers/credentials';

const authOptions = {
	providers: [
		CredentialsProvider({
			name: "credentials",
			credentials: {
				email: {
					label: "email",
					type: "email",
					placeholder: "john@doe.com.sg",
				},
				password: { label: "password", type: "password" },
			},
			authorize: async (credentials) => {
				const request = axios.post("https://itsag1t2.com/auth/login", {
					email: credentials.email,
					password: credentials.password,
				});

				return await request
					.then((response) => {
						const user = {
							user_id: response.data.data.user_id,
							token: "Bearer " + response.data.data.token,
							is_new: response.data.data.is_new,
						};
						return user;
					})
					.catch((error) => {
						console.log(error);
						return null;
					});
			},
		}),
	],
	callbacks: {
		jwt: async ({ token, user }) => {
			if (user) {
				token.id = user.token;
				token.userId = user.user_id;
				token.is_new = user.is_new; // Add isFirstTimeLogin to user object
			}

			return token;
		},
		session: ({ session, token }) => {
			if (token) {
				session.id = token.id;
				session.userId = token.userId;
				session.is_new = token.is_new; // Add isFirstTimeLogin to user object
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
