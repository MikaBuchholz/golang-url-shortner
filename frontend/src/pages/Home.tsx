import { Component, createSignal } from "solid-js";

export interface UrlModel {
  id: string;
  view_count: number;
}

const Home: Component = () => {
  const [input, setInput] = createSignal("");
  const [shortenedUrlModel, setshortenedUrlModel] = createSignal<UrlModel>();

  return (
    <>
      <div class="flex justify-center">
        <div class="w-[30vw]">
          <form
            onSubmit={async (e) => {
              e.preventDefault();
              //`{\"payload\": \"${input()}\"}`
              const result = await fetch(
                "http://localhost:8080/api/v1/url/new",
                {
                  body: `{\"payload\": \"${input()}\"}`,
                  method: "POST",
                }
              );

              let data = (await result.json()) as UrlModel;
              setshortenedUrlModel(data);
            }}
          >
            <label
              for="url"
              class="block mb-2 text-sm font-medium text-black-900 dark:text-white"
            >
              Url
            </label>

            <input
              placeholder="Your URL..."
              id="url"
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              value={input()}
              type="text"
              onChange={(e) => {
                setInput(e.target.value);
              }}
            />
          </form>
        </div>
      </div>
      {!shortenedUrlModel() || !shortenedUrlModel()?.id.length ? null : (
        <>
          <div class="flex justify-center cursor-pointer">
            <div class="rounded p-3 mt-10 bg-sky-700 hover:bg-sky-400">
              <a
                class="text-white"
                href={`${window.location.origin}/${shortenedUrlModel()?.id}`}
              >{`${window.location.origin}/${shortenedUrlModel()?.id}`}</a>
            </div>
          </div>
        </>
      )}
    </>
  );
};

export default Home;
