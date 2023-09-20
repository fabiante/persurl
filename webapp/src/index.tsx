import { render } from 'preact';
import { LocationProvider, Router, Route } from 'preact-iso';

import { Header } from './components/Header.jsx';
import { Home } from './pages/Home';
import { NotFound } from './pages/_404.jsx';
import './style.css';
import {route} from "./route";

export function App() {
	return (
		<LocationProvider>
			{/*<Header />*/}
			<main>
				<Router>
					<Route path={route("/")} component={Home} />
					<Route default component={NotFound} />
				</Router>
			</main>
		</LocationProvider>
	);
}

render(<App />, document.getElementById('app'));
