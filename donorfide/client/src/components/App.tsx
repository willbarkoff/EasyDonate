import * as React from "react";

import "./App.sass";
import * as api from "../api";
import * as stripe from "../stripe";

import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Nav from "./Nav";
import Footer from "./Footer";
import LoadingScreen from "./LoadingScreen";
import NotFound from "./pages/notfound/NotFound";
import DonatePage from "./pages/donate/DonatePage";
import DonateCreditInfo from "./pages/donate/DonateCreditInfo";

import "./ui/stripe.sass";
import { Elements } from "@stripe/react-stripe-js";
import DonationSuccess from "./pages/donate/DonationSuccess";
import AdminPage from "./pages/admin/AdminPage";
import LoginPage from "./pages/LoginPage";
import RequireLogin from "./pages/RequireLogin";
import AdminSettingsPage from "./pages/admin/AdminSettingsPage";

function useOrgContext() {
	const [isLoading, setIsLoading] = React.useState(true);
	const [contextData, setContextData] = React.useState(null as unknown as api.contextOrg);

	React.useEffect(() => {
		async function fetchContext() {
			const data = await api.GET<api.contextOrg>("context/org");
			setContextData(data);
			document.title = "Donate | " + data.name;
			await stripe.prepareStripe(data.stripe_pk);
			setIsLoading(false);
		}

		fetchContext();
	}, []);

	return [contextData, isLoading];
}

export const OrgContext = React.createContext(null as unknown as api.contextOrg);

export default function App(): JSX.Element {
	const [contextData, isLoading] = useOrgContext();

	if (isLoading) {
		return <LoadingScreen />;
	}

	return <OrgContext.Provider value={contextData as api.contextOrg}>
		<Router>
			<Elements stripe={stripe.stripe}>
				<Nav />
				<Switch>
					<Route path="/" exact><DonatePage /></Route>
					<Route path="/donation/success"><DonationSuccess /></Route>
					<Route path="/donate/:amount" exact><DonateCreditInfo /></Route>
					<Route path="/login"><LoginPage /></Route>
					<Route path="/admin/settings"><RequireLogin component={() => <AdminSettingsPage />} /></Route>
					<Route path="/admin"><RequireLogin component={me => <AdminPage me={me} />} /></Route>
					<Route path="/"><NotFound /></Route>
				</Switch>
				<Footer />
			</Elements>
		</Router>
	</OrgContext.Provider>;
}