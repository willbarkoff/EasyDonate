import * as React from "react";
import * as api from "../../api";
import LoadingScreen from "../LoadingScreen";
import {Redirect, useLocation} from "react-router-dom";

interface requireLoginProps {
	component(me: api.me): JSX.Element
}

export default function RequireLogin(props: requireLoginProps) {
	const [isLoading, setIsLoading] = React.useState(true)
	const [me, setMe] = React.useState(null as null | api.me)
	const location = useLocation()

	React.useEffect(() => {
		async function fetchMe() {
			try {
				let data = await api.GET<api.me>("auth/me")
				setMe(data)
			} catch (e) {
			}
			setIsLoading(false)
		}

		fetchMe()
	}, [])

	if (isLoading) {
		return <LoadingScreen/>
	}

	if (me == null) {
		return <Redirect to={`/login?done=${encodeURIComponent(location.pathname)}`}/>
	}

	return props.component(me)
}