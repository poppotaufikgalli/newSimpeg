export default defineEventHandler(async(event) => {
    const config = useRuntimeConfig()
    const ep = event.context.params._
    const {token} = parseCookies(event)

    const form = await readMultipartFormData(event)
    const formdata = new FormData();

    if(form){
        //formdata.append(form[0].name, new Blob([form[1].data, {type: form[1].type}]))
        for (const i in form) {
            console.log(i, form[i].name, form[i])
            if(form[i].name == 'file'){
                formdata.append(form[i].name, new Blob([form[i].data, {type: form[i].type}]), form[i].filename);
            }else{
                formdata.append(form[i].name, form[i].data);
            }
        }    
    }

    return await fetch(config.public.apiBase+"/"+ep, {
        method: 'POST',
        headers: {
            "Authorization": "Bearer " + token,
        },
        body: formdata,
    }).then((response) => 
        response.json().then(data => ({status: response.status, data: data}))
    )
})