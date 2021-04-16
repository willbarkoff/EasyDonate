import * as React from "react";
import * as api from "../../../api"

interface adminStringSettingProps {
	originalValue: string
	settingKey: string
	label: string
	description: string
}

export default function AdminStringSetting(props: adminStringSettingProps) {
	let [value, setValue] = React.useState(props.originalValue)
	let [error, setError] = React.useState("")
	let [success, setSuccess] = React.useState(false)
	let [loading, setLoading] = React.useState(false)

	console.log(props)

	async function update() {
		try {
			setLoading(true)
			await api.POST<undefined>("admin/updateSetting", {key: props.settingKey, value})
			setSuccess(true)
		} catch (e) {
			setError(api.errorToEnglish(e.message))
		} finally {
			setLoading(false)
		}
	}

	return <div className="field">
		<label className="label">{props.label}</label>
		<div className="field has-addons">
			<div className="control is-expanded">
				<input className={`input ${error ? "is-danger" : ""} ${success ? "is-success" : ""}`} type="text"
					   placeholder={props.originalValue} value={value}
					   onChange={e => {
						   setValue(e.target.value);
						   setSuccess(false)
						   setError("")
					   }}
					   disabled={loading}/>
			</div>
			<div className="control">
				<button className={`button is-info ${loading ? "is-loading" : ""}`} onClick={update} disabled={loading}>
					Save
				</button>
			</div>
			<div className="control">
				<button className="button is-danger" onClick={() => setValue(props.originalValue)} disabled={loading}>
					Reset
				</button>
			</div>
		</div>
		<p className={`help ${error ? "is-danger" : ""}`}>{error || props.description}</p>
	</div>
}