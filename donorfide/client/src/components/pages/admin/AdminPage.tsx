import * as React from "react";
import Hero from "../../ui/Hero";
import Page from "../Page";
import * as api from "../../../api";
import { Link } from "react-router-dom";

interface adminPageProps {
	me: api.me
}

export default function AdminPage(props: adminPageProps): JSX.Element {

	return <Page>
		<Hero title="Administration" subtitle="Manage donorfide" center />
		<section className="section">
			<div className="container">
				<p className="block">Hello, <strong>{props.me.first_name}</strong>.</p>
				{props.me.level < 1 ?
					<div className="notification is-danger is-light">You do not have permission to
						access this page.</div>
					: <div>
						<strong className="block">Administration tools</strong>
						<ul>
							{props.me.level >= 2 && <li><Link to="/admin/settings">Settings</Link></li>}
							{props.me.level >= 2 && <li><Link to="/admin/users">Users</Link></li>}
							<li><Link to="/admin/donations">Donations</Link></li>
							<li><a href="https://donorfide.org/docs">Donorfide Documentation</a></li>
						</ul>
					</div>
				}
			</div>
		</section>
	</Page>;
}