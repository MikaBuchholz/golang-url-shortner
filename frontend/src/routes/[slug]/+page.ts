export const load = async ({ params }) => {
    const response = await fetch(`http://localhost:8080/api/v1/url/${params.slug}`, {
        method: "GET"
    })

    return {
        response: await response.json() as ApiGetResponse
    }
}