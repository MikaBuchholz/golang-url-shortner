<script lang="ts">
	import { goto } from '$app/navigation';
	import { error } from '@sveltejs/kit';
	import { page } from '$app/stores';

	let inputState = '';
	$: shortenedUrl = '';

	const shortenURL = async () => {
		const body = `{"payload":"${inputState.trim()}"}`;

		const response = await fetch(`http://localhost:8080/api/v1/url/new`, {
			method: 'POST',
			body
		}).catch(() => {
			throw error(503, { message: 'Service currently not working :|' });
		});

		let apiPostResponse = (await response.json()) as ApiPostResponse;

		console.log(apiPostResponse);

		return window.location.toString() + apiPostResponse.id;
	};
</script>

<div class="h-screen bg-slate-800">
	<div class="flex justify-center">
		<div class="flex flex-row gap-6 max-h-3.5 mt-20 items-center">
			<div class="inline-block w-[20vw]">
				<input
					bind:value={inputState}
					type="url"
					id="website"
					class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					placeholder="youtube.com"
					required
				/>
			</div>
			<button
				on:click={async () => {
					shortenedUrl = await shortenURL();
				}}
				class="inline-flex items-center gap-2 rounded border border-indigo-600 bg-indigo-600 px-3 py-2 text-white hover:bg-indigo-500 hover:text-white-200 focus:outline-none focus:ring active:text-indigo-500"
			>
				<span class="text-sm font-small"> Shorten </span>
				<svg
					class="h-5 w-5 rtl:rotate-180"
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M17 8l4 4m0 0l-4 4m4-4H3"
					/>
				</svg>
			</button>
		</div>
	</div>
	<div class="flex justify-center">
		<div class="flex flex-row justify-evenly">
			{#if shortenedUrl.length}
				<div class="inline-block mt-10">
					<button
						on:click={() => {
							goto(shortenedUrl);
						}}
						class=" mt-22 bg-green-100 text-green-800 text-md font-medium mr-2 px-2.5 py-0.5 rounded dark:bg-green-900 dark:text-green-300"
						>{shortenedUrl}</button
					>

					<button
						on:click={() => {
							navigator.clipboard.writeText(shortenedUrl);
						}}
						class="inline-flex items-center gap-2 rounded border border-slate-600 bg-slate-800 px-3 py-2 text-white hover:bg-slate-600 hover:text-white-200 focus:outline-none"
					>
						<svg
							class="w-4 h-4 stroke-current"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
							><path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M9 12h3.75M9 15h3.75M9 18h3.75m3 .75H18a2.25 2.25 0 002.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 00-1.123-.08m-5.801 0c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75 2.25 2.25 0 00-.1-.664m-5.8 0A2.251 2.251 0 0113.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m0 0H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V9.375c0-.621-.504-1.125-1.125-1.125H8.25zM6.75 12h.008v.008H6.75V12zm0 3h.008v.008H6.75V15zm0 3h.008v.008H6.75V18z"
							/></svg
						>
					</button>
				</div>
			{/if}
		</div>
	</div>
</div>
