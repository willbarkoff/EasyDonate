import * as React from "react";

import "./App.sass"
import * as api from "../api"
import * as stripe from "../stripe"

import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import Nav from "./Nav";
import Footer from "./Footer";
import LoadingScreen from "./LoadingScreen";
import NotFound from "./pages/notfound/NotFound";
import DonatePage from "./pages/donate/DonatePage";
import DonateCreditInfo from "./pages/donate/DonateCreditInfo";

import "./ui/stripe.sass"

function useOrgContext() {
	const [isLoading, setIsLoading] = React.useState(true);
	const [contextData, setContextData] = React.useState(null as unknown as api.contextOrg);

	React.useEffect(() => {
		async function fetchContext() {
			let data = await api.GET<api.contextOrg>("context/org")
			await setContextData(data)
			await stripe.prepareStripe(data.stripe_pk)
			setIsLoading(false)
		}

		fetchContext()
	}, []);

	return [contextData, isLoading];
}

export const OrgContext = React.createContext(null as unknown as api.contextOrg)

export default function App() {
	const [contextData, isLoading] = useOrgContext()

	if (isLoading) {
		return <LoadingScreen/>
	}

	return <OrgContext.Provider value={contextData as api.contextOrg}>
		<Router>
			<Nav/>
			<Switch>
				<Route path="/" exact><DonatePage/></Route>
				<Route path="/donate/:amount" exact><DonateCreditInfo/></Route>
				<Route path="/"><NotFound/></Route>
			</Switch>
			<Footer/>
		</Router>
	</OrgContext.Provider>
}