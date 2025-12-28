const getCanvasHash = async (canvas) => {
    const dataURL = canvas.toDataURL('image/png');

    // Convert the string to a Uint8Array
    const msgUint8 = new TextEncoder().encode(dataURL);

    // Hash the data using SHA-256
    const hashBuffer = await crypto.subtle.digest('SHA-256', msgUint8);

    // Convert the ArrayBuffer to a hexadecimal string
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');

    return ['canvas_hash', hashHex];
};

const PROD_URL = 'https://track.danielberry.tech/stat';
const DEV_URL = '/stat';


const onload = () => {
	let payload = {
		timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
		screen_width: screen.width,
		screen_height: screen.height,
		num_cpu_cores: navigator.hardwareConcurrency,
		language: navigator.language,
		user_agent: navigator.userAgent,
		//canvas_hash: 
	};
	Promise.allSettled([
		getCanvasHash(document.createElement('canvas')),
	])
		.then((results) => {
			results.forEach(({status, value}) => {
				if(status !== 'fulfilled') {
					return;
				}
				let [tag, val] = value;
				payload[tag] = val;
			});

			fetch(DEV_URL, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(payload),
			});
		});
};

window.addEventListener("load", onload);
