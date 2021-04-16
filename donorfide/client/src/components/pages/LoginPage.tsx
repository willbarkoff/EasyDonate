import * as React from 'react'
import {FormEvent} from 'react'
import Page from "./Page";
import Hero from "../ui/Hero";
import {Redirect, useLocation} from "react-router-dom";
import * as api from "../../api"

function useQuery() {
	return new URLSearchParams(useLocation().search);
}

export default function LoginPage() {
	let query = useQuery()
	let done = query.get("done") || "/"

	const [email, setEmail] = React.useState("")
	const [password, setPassword] = React.useState("")
	const [isLoading, setIsLoading] = React.useState(false)
	const [error, setError] = React.useState("" as any)
	const [isDone, setIsDone] = React.useState(false)

	const submit = async (event: FormEvent) => {
		event.preventDefault();
		setIsLoading(true)
		try {
			await api.POST("auth/login", {email: email, password: password})
			setIsDone(true)
		} catch (e) {
			setIsLoading(false)
			setError(e as string)
		}
	}

	if (isDone) {
		return <Redirect to={done}/>
	}

	return <Page>
		<Hero title="Log in"/>
		<section className="section">
			<div className="container">
				{error && <div className="notification is-danger is-light">{api.errorToEnglish(error.message)}</div>}
				{done != "/" &&
                <div className="notification is-info is-light">You need to log in to access that page.</div>}
				<form className="form block" onSubmit={submit}>
					<div className="field">
						<label className="label">Email</label>
						<input className="input" type="email" placeholder="jane@appleseed.com" required
							   disabled={isLoading} value={email} onChange={e => setEmail(e.target.value)}/>
					</div>
					<div className="field">
						<label className="label">Password</label>
						<input className="input" type="password" required disabled={isLoading} value={password}
							   onChange={e => setPassword(e.target.value)}/>
					</div>
					<input type="submit" className={`button is-primary ${isLoading ? "is-loading" : ""}`}
						   value="Log in!" disabled={isLoading}/>
				</form>
			</div>
		</section>
	</Page>
}