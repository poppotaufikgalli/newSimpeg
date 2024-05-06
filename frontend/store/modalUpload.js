export const useModalUpload = defineStore('modal-upload', () => {
    const displayModal = ref(false)
    const modalTitle = ref('')
    const modalText = ref('')
    const modalAction = ref(false)
    const modalFile = ref("https://upload.wikimedia.org/wikipedia/commons/6/65/No-Image-Placeholder.svg")
    const modalFileType = ref('')
    const fileBlob = ref(null)

    function showModal(title, typeFile, file) {
        modalTitle.value = title
        modalFileType.value = typeFile
        modalFile.value = file ? file : modalFile.value
        displayModal.value = true
    }

    function closeModal(){
        modalAction.value = false
        displayModal.value = false
    }

    function doAction() {
        modalAction.value = true
        displayModal.value = false
    }

    return { 
        displayModal, 
        modalTitle, 
        modalFile, 
        modalFileType,
        showModal, 
        closeModal, 
        doAction, 
        fileBlob,
    }
})