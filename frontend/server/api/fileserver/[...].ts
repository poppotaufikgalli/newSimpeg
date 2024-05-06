export default defineEventHandler( async (e) => {
    const ep = e.context.params._
    const {token} = parseCookies(e)
    const config = useRuntimeConfig()

    try {
        const blob = await $fetch(`${config.public.apiBase}/${ep}`, {
            headers: {
                "Authorization": "Bearer " + token,
            },
        });

        return blob
    }
    catch (err) {
        console.error(err);
    }

    return null;
});