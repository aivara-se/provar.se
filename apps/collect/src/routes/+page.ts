import { goto } from '$app/navigation';
import { browser } from '$app/environment';
import { getConfig } from '$lib/client/store';

function redirect() {
	if (getConfig().configured) {
		goto('/admin');
	} else {
		goto('/setup');
	}
}

if (browser) {
	redirect();
}
