<script setup>
import { useModalUploadDoc } from '~/store/modalUploadDoc';
import { useAuthStore } from '~/store/auth'; // import the auth store we just created

const modalUploadDoc = useModalUploadDoc();
const { fileBlob } = storeToRefs(modalUploadDoc)
const { userInfo } = useAuthStore();
const { $decodeBase64, $encodeBase64 } = useNuxtApp()
const route = useRoute()
const dayjs = useDayjs()
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat)

const {snip, stmtjab, stmtMutasi, sid} = route.params
const {sbkn } = route.query

/*const dataRefJabatan = ref({
	nip: '',
	tmtjab: '',
	tmt_mutasi: '',
})*/

const dataJabatan = ref({
	nip: '',
	kpej: '',
	nskjabat:'',
	tskjabat: '',
	tmtjab: '',
	tlantik: '',
	id_opd: '',
	id_sub_opd: '',
	id_unor_bkn: '',
	jnsjab: '',
	keselon: 0,
	sub_jab:'',
	jns_tugas_mutasi: '',
	tmt_mutasi: '',
	kjab: 0,
	njab: '',
	njab0: '',
	akhir: 0,
	nunker: '',
	idSync:'',
	filename: '',
	filename_jab: '',
	filename_lantik: '',
	filename_tugas_mutasi: '',
})

const reqData = ref([
	'kpej', 'tskjabat', 'tmtjab', 'tlantik', 'id_opd', 'id_sub_opd', 'id_unor_bkn', 'jnsjab', 'keselon', 'jns_tugas_mutasi', 'tmt_mutasi', 'kjab',
])

const refUpdateBkn = ref({
	"eselonId": "keselon",
	"id": "idSync",
	"instansiId": "id_instansi",
	//"jabatanFungsionalId": "kjab",
	//"jabatanFungsionalUmumId": "kjab",
	"jenisJabatan": "jnsjab",
	//"jenisMutasiId": "jns_tugas_mutasi",
	//"jenisPenugasanId": "jns_tugas_mutasi",
	"nomorSk": "nskjabat",
	/*"path": [
		{
			"dok_id": "string",
			"dok_nama": "string",
			"dok_uri": "string",
			"object": "string",
			"slug": "string"
		}
	],*/
	//"pnsId": "id_ref_bkn",
	//"satuanKerjaId": "id_instansi",
	"subJabatanId": "sub_jab",
	"tanggalSk": "tskjabat",
	"tmtJabatan": "tmtjab",
	"tmtMutasi": "tmt_mutasi",
	"tmtPelantikan": "tlantik",
	"unorId": "id_unor_bkn"
})

const pejabats = ref([])
const jenis_jabatans = ref([])
const eselons = ref([])
const opds = ref([])

const jabatans = ref([])
const apiUrl = ref(null) 

const refJnsTugasMutasi = ref();

const selopd = ref('')
const selkjab = ref('')
const selsubjab = ref('')
const jnstugasmutasis = ref()

const subjabs = ref([])

const refPejabat = ref(null)
const refOpd = ref(null)
const refSubOpd = ref(null)
const refJnsJab = ref(null)
const refEselon = ref(null)
const refJabatan = ref(null)
const refSubJab = ref(null)
const method = ref('Simpan')

const { pending: pendingRef, data: dataRef, refresh: refreshDataRef } = await useLazyAsyncData('getDataRef', async () => {
	//loading.value = true
		const [pejabat, jenis_jabatan, eselon, opd, jnstugasmutasi] = await Promise.all([
			$fetch('/api/gets/pejabat'),
			$fetch('/api/gets/jenis_jabatan'),
			$fetch('/api/gets/eselon'),
			$fetch('/api/posts/opd', {
				method: 'POST',
				body: JSON.stringify({
					sfilter: [1,0]
				})
			}),
			$fetch('/api/gets/jenis_tugas_mutasi'),
		])

		pejabats.value = pejabat
		opds.value = opd
		jenis_jabatans.value = jenis_jabatan
		eselons.value = eselon
		jnstugasmutasis.value = jnstugasmutasi
		jenis_tugas_mutasi.value = jnstugasmutasi

		if(dataJabatan.value.jnsjab != ""){
			jenis_tugas_mutasi.value = _filter(jnstugasmutasi, function(o){
				var jns = o.jns.split(',')
				return _includes(jns, dataJabatan.value.jnsjab)
			})
		}


	if(refOpd){
		refOpd.value.reinit()
	}

	if(refJnsTugasMutasi){
		refJnsTugasMutasi.value.reinit()
	}

	if(refJnsJab){
		refJnsJab.value.reinit()
	}

	if(refEselon){
		refEselon.value.reinit()
	}

	if(refPejabat){
		refPejabat.value.reinit()
	}

});

const { pending, data, refresh} = await useAsyncData('getDataJabatan', async() => {
	if(snip){
		let nip = $decodeBase64(snip)
		if(sbkn){
			let sid = $decodeBase64(sbkn)
			var result = await $fetch(`/api/gets/singkronisasi_bkn/getDataBKNById/jabatan/${sid}`)
			var item = result.data

			console.log(item)
			let tsk = dayjs(dayjs(item.tanggalSk, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
			let tmtjab = dayjs(dayjs(item.tmtJabatan, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
			let tlantik = dayjs(dayjs(item.tmtPelantikan, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")
			let tmt_mutasi = dayjs(dayjs(item.tmtMutasi, "DD-MM-YYYY")).format("YYYY-MM-DD[T]HH:mm:ss[Z]")

			var result_opd = await $fetch("/api/posts/opd/", {
				method: "POST",
				body: JSON.stringify({
					id_unor_bkn: item.unorId
				})
			})

			var jns_tugas_mutasi = item.jenisPenugasanId != "" ? item.jenisPenugasanId : item.jenisMutasiId
			var kjab = item.jenisJabatan == "2" ? item.jabatanFungsionalId : (item.jenisJabatan == "4" ? item.jabatanFungsionalUmumId : "")  
			var njab0 = item.jenisJabatan == "2" ? item.jabatanFungsionalNama : (item.jenisJabatan == "4" ? item.jabatanFungsionalUmumNama : "")  
			var njab = njab0.toUpperCase() + " PADA " +  item.unorNama + " " + item.unorIndukNama

			dataJabatan.value = {
				nip: nip,
				nskjabat: item.nomorSk,
				tskjabat: tsk,
				tmtjab: tmtjab,
				tlantik: tlantik,
				id_opd: result_opd[0].parent_opd,
				id_sub_opd: result_opd[0].id,
				id_unor_bkn: item.unorId,
				jnsjab: item.jenisJabatan,
				keselon: item.eselonId,
				sub_jab:item.subJabatanId,
				jns_tugas_mutasi: jns_tugas_mutasi,
				tmt_mutasi: tmt_mutasi,
				kjab: kjab,
				kjab_bkn: kjab,
				njab: njab,
				njab0: njab0,
				nunker: item.unorNama,
				idSync:item.id,
				akhir: 0,
			}

			selopd.value = result_opd[0].id
			
			if(refOpd){
				refOpd.value.reinit()
			}

			selkjab.value = kjab

			/*if(result.jns_tugas_mutasi == 'MU'){
				//console.log("-->")
				await getRiwayatJab()
				//console.log("<--")
			}*/

			if(item.jenisJabatan == "2"){
				getSubJab(result.kjab)
				selsubjab.value = result.sub_jab
			}
			
			console.log(dataJabatan.value)

		}else{
			if(route.query.akhir && (route.query.akhir == "true" || route.query.akhir == true)){
				var result = await $fetch('/api/gets/riwayat_jabatan/'+nip+'/akhir');

				if(result.nip == nip){
					method.value = "Update"

					dataJabatan.value = _pickBy(result, function(val, key) {
						return _includes(_keys(dataJabatan.value), key);
					})

					selopd.value = result.id_sub_opd
					selkjab.value = result.kjab
					
					CariJabatanNonFormasi()
				}
				akhir.value = route.query.akhir ? 1 : 0
			}else{
				let id = $decodeBase64(sid)
				var result = await $fetch('/api/gets/riwayat_jabatan/'+nip+'/'+id);			

				if(result.nip == nip){
					method.value = "Update"

					dataJabatan.value = _pickBy(result, function(val, key) {
						return _includes(_keys(dataJabatan.value), key);
					})

					/*dataRefJabatan.value = _pickBy(result, function(val, key) {
						return _includes(_keys(dataRefJabatan.value), key);
					})*/

					selopd.value = result.id_sub_opd

					if(dataJabatan.value.id_unor_bkn == null || dataJabatan.value.id_unor_bkn == ''){
						getUnorBkn(dataJabatan.value.id_sub_opd)
					}
					
					selkjab.value = result.kjab
					//jnsjab.value = result.jnsjab

					if(result.jns_tugas_mutasi == 'MU'){
						//console.log("-->")
						await getRiwayatJab()
						//console.log("<--")
					}

					if(result.jnsjab == 2){
						getSubJab(result.kjab)
						selsubjab.value = result.sub_jab
					}

					console.log("load data 2")
				}
			}
		}
	}
})

const jenis_tugas_mutasi = ref([])

const tskjabat = computed({
	get(){
		return dayjs(dataJabatan.value.tskjabat).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataJabatan.value.tskjabat = val
	}
})

const tmtjab = computed({
	get(){
		return dayjs(dataJabatan.value.tmtjab).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataJabatan.value.tmtjab = val
	}
})

const tlantik = computed({
	get(){
		return dayjs(dataJabatan.value.tlantik).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataJabatan.value.tlantik = val
	}
})

const tmt_mutasi = computed({
	get(){
		return dayjs(dataJabatan.value.tmt_mutasi).format("YYYY-MM-DD").toString()
	},
	set(val){
		dataJabatan.value.tmt_mutasi = val
	}
})

const jnsjab = computed({
	get(){
		return dataJabatan.value.jnsjab
	},
	set(val){
		dataJabatan.value.jnsjab = val
		//console.log("change jnsjab")

		jenis_tugas_mutasi.value = _filter(jnstugasmutasis.value, function(o){
			var jns = o.jns.split(',')
			return _includes(jns, val)
		})

		if(refJnsTugasMutasi){
			refJnsTugasMutasi.value.reinit()
		}

		CariJabatanNonFormasi()
	}
})

watch(selopd, (value) => {
	console.log("change selopd")
	dataJabatan.value.id_opd = value.substring(0,3)
	dataJabatan.value.id_sub_opd = value

	getUnorBkn(value)

	if(dataJabatan.value.jns_tugas_mutasi == 'MU'){
		getRiwayatJab()
	}

	//console.log("load data selopd")
	CariJabatanNonFormasi()
})

async function getUnorBkn(id){
	var result = await $fetch("/api/gets/opd/"+id)

	if(result.id_unor_bkn == '' || result.id_unor_bkn == undefined || result.id_unor_bkn == null){
		toast.add({
			id: 'error_check_unor_bkn',
			icon: "i-heroicons-x-circle",
			description: `OPD yang dipilih belum memiliki Unor BKN`,
			timeout: 6000,
			color: 'red',
		}) 		
	}else{
		dataJabatan.value.id_unor_bkn = result.id_unor_bkn
	}
}

async function getRiwayatJab(){
	console.log("getRiwayatJab")
	let nip = $decodeBase64(snip)
	var jnsjab = dataJabatan.value.jnsjab
	var result = await $fetch("/api/posts/riwayat_jabatan/", {
		method: 'POST',
		body: JSON.stringify({
			nip: nip,
			jnsjab: jnsjab
		})
	})

	if(result){
		jabatans.value = result.map(function(o) {
			//var kjab = o.kjab+":"+o.kjab_bkn+":"+dayjs(dayjs(o.tmtjab, "YYYY-MM-DD")).format("YYYY-MM-DD")
			var tmtjab0 = dayjs(dayjs(o.tmtjab, "YYYY-MM-DD")).format("YYYY-MM-DD")
			//var kjab0 = o.kjab+":"+o.kjab_bkn+":"+tmtjab0
			var kjab0 = o.kjab
			var njab0 = o.njab.split("PADA")
				return _extend({}, o, {
					keyList: kjab0, 
					namaList: (o.njab0 ?? njab0[0]) + " / "+tmtjab0,
				})
		});

		console.log(jabatans.value)

		if(refJabatan){
			refJabatan.value.reinit()
		}
	}
}

async function getSubJab(jft){
	var result = await $fetch("/api/gets/jabatan_sub_bkn/filter_jft/"+jft)

	if(result && result.length > 0){
		subjabs.value = result

		if(refSubJab){
			refSubJab.value.reinit()
		}
	}
}

watch(selkjab, (value) => {
	console.log("selkjab change", value)
	var jab = refJabatan.value.readOptions(value)
	
	dataJabatan.value.kjab = value
	if(dataJabatan.value.jns_tugas_mutasi == 'MU'){
		var njab = jab.split(" / ")
		jab = (njab[0].split("-"))[1]
		console.log(njab)
		var dtJab = _find(jabatans.value, function(o){
			var tmtjab0 = dayjs(dayjs(o.tmtjab, "YYYY-MM-DD")).format("YYYY-MM-DD")
			return o.kjab == value && tmtjab0 == njab[1]
		})

		dataJabatan.value.kjab_bkn = dtJab.kjab_bkn
		dataJabatan.value.tmtjab = dtJab.tmtjab
	}else{
		dataJabatan.value.kjab_bkn = value	
	}
	
	dataJabatan.value.njab0 = jab.toUpperCase();	
	dataJabatan.value.njab = dataJabatan.value.njab0.toUpperCase()+" PADA "+dataJabatan.value.nunker;
})

function stackOPD (opd, njab) {
	var parent_opd_id = opd.parent_opd 
	while(parent_opd_id != null){
		var result = _find(opds.value, function(o){
			return o.id == parent_opd_id
		})	

		if(result.id != 1){
			njab.push(result.nama)	
		}else{
			njab.push("KOTA TANJUNGPINANG")	
		}

		parent_opd_id = result.parent_opd
	}

	return njab
}

const CariJabatanNonFormasi = () => {
	var njab = []
	var result = _find(opds.value, {id: selopd.value})

	if(result){
		if(dataJabatan.value.jnsjab == 1){
			//console.log("cari jabatan Struktural")

			dataJabatan.value.kjab = selopd.value
			dataJabatan.value.kjab_bkn = result.id_unor_bkn
			dataJabatan.value.keselon = result.id_eselon

			if(result.nama_jabatan != "" && result.nama_jabatan != null){
				if(result.nama != result.nama_jabatan){
					njab = [result.nama_jabatan.toUpperCase(), "PADA", result.nama]	
				}else{
					njab = [result.nama_jabatan.toUpperCase(), "PADA"]	
				}

				dataJabatan.value.njab0 = result.nama_jabatan
			}else{
				njab = [result.nama.toUpperCase(), "PADA"]	
				dataJabatan.value.njab0 = result.nama
			}

			//console.log(dataJabatan.value)

		}else{
			//console.log("cari jabatan Non Struktural")

			njab = [result.nama]
			dataJabatan.value.keselon = '99'
		}

		var newJab = stackOPD(result, njab)

		var nunker = [result.nama]
		var newNunker = stackOPD(result, nunker)
		dataJabatan.value.nunker = newNunker.join(" ")

		dataJabatan.value.njab = newJab.join(" ")

		if(refEselon){
			refEselon.value.reinit()
		}
	}
}

onMounted(() => {
	refreshNuxtData(["getDataRef", "getDataJabatan"])
	modalUploadDoc.$onAction(callback, false)
})

const showSpinner = ref(false)
const toast = useToast()

async function simpanData(nMethod) {
	showSpinner.value = true
	let canUpdateBkn = false

	//console.log(dataJabatan.value)
	var compacted = _pickBy(dataJabatan.value);
	compacted.tmtjab = compacted.tmtjab ? dayjs(dayjs(compacted.tmtjab, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tskjabat = compacted.tskjabat ? dayjs(dayjs(compacted.tskjabat, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.tlantik = compacted.tlantik ? dayjs(dayjs(compacted.tlantik, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	//compacted.tmt_mutasi = compacted.jns_tugas_mutasi == 'MJ' ? compacted.tmtjab :  compacted.tmt_mutasi
	compacted.tmt_mutasi = compacted.tmt_mutasi ? dayjs(dayjs(compacted.tmt_mutasi, "YYYY-MM-DD")).format("YYYY-MM-DD[T]HH:mm:ss[Z]") : null
	compacted.keselon = compacted.keselon *1
	//compacted.status = compacted.status *1
	compacted.kpej = compacted.kpej *1
	compacted.id_instansi = "A5EB03E23CFBF6A0E040640A040252AD"
	compacted.nama_instansi = 'PEMERINTAH KOTA TANJUNGPINANG'
	compacted.jns_tugas_mutasi = compacted.jnsjab == 1 ? 'D' : compacted.jns_tugas_mutasi
	compacted.sub_jab = selsubjab.value
	compacted.akhir = compacted.akhir == true ? 1 : 0

	//var dr = dataRefJabatan.value
	//dr.kgolru = dr.kgolru *1
	//console.log(compacted)
	let id = $decodeBase64(sid)
	
	var errorList = []

	for (var i = reqData.value.length - 1; i >= 0; i--) {
		var idx = reqData.value[i];
		if(compacted[idx] == undefined){
			if(idx == 'tmt_mutasi' && _includes(['MJ','D','PJ','PLH', 'TT'],compacted.jns_tugas_mutasi)){

			}else{
				errorList.push(idx)
			}
		}
	}

	if(errorList.length > 0){
		toast.add({
			id: 'error_check_jabatan',
			icon: "i-heroicons-question-mark-circle",
			title: "Validasi Data BKN",
			description: `data ${errorList.join(", ")} tidak boleh kosong`,
			timeout: 6000,
			color: 'orange',
		})
		showSpinner.value = false
		return
	}

	if(nMethod == 'Update'){
		let nip = $decodeBase64(snip)
		var body = {
			//"refdata" : dr,
			refdata:{
				nip: nip,
				id: id*1,
			},
			"data" : compacted
		}

		var {data, error} = await useFetch('/api/puts/riwayat_jabatan', {
			method: 'PUT',
			body: JSON.stringify(body),
		})

		if(error.value){
			const {title, description} = error.value.data
			toast.add({
				id: 'error_put_jabatan',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_put_jabatan',
				icon: "i-heroicons-check-circle",
				title: "Update Data Berhasil",
				description: "Data telah di Update",
				timeout: 6000
			}) 	
			sendDataBKN(JSON.stringify(compacted))
		  	navigateTo({ path: '/pegawai/'+snip+'/RiwayatPegawai/jabatan/'})
		}
	}else{
		compacted.nip = $decodeBase64(snip)
		//compacted.akhir = 1
		var body = JSON.stringify(compacted)

		var {data, error} = await useFetch('/api/posts/riwayat_jabatan/new', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			const {title, description} = error.value.data.data
			toast.add({
				id: 'error_post_jabatan',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_post_jabatan',
				icon: "i-heroicons-check-circle",
				description: "Data berhasil di Ditambahkan",
				timeout: 6000
			}) 	

			compacted.id = data.value.id
			sendDataBKN(JSON.stringify(compacted))
		  	navigateTo({ path: '/pegawai/'+snip+'/RiwayatPegawai/jabatan/'})
		}
	}

	showSpinner.value = false
	//refreshNuxtData(["getDataRef", "getDataJabatan"])
}

async function sendDataBKN(body){
	if(userInfo.gid == 1){
		var {data, error} = await useFetch('/api/posts/riwayat_jabatan/bkn', {
			method: 'POST',
			body: body,
		})

		if(error.value){
			const {title, description} = error.value.data.data
			toast.add({
				id: 'error_post_jabatan',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red',
			}) 	
		}else{
			toast.add({
				id: 'success_post_jabatan',
				icon: "i-heroicons-check-circle",
				title: "Update Data Jabatan BKN",
				description: "Data berhasil di Update",
				timeout: 6000
			}) 	
		}	
	}
}

const akhir = computed({
	get(){
		return dataJabatan.value.akhir == 0 ? false : true
	},
	set(val){
		dataJabatan.value.akhir = val
	}
})

const aakhir = computed({
	get(){
		return dataJabatan.value.akhir == 0 ? 'Tidak' : 'Ya'
	},
})

const selUploadDoc = ref("")
const selUploadDocJudul= ref("")
const selUploadDocBkn = ref("")

const doUploadDoc = async(judul, fileIdx, idxBkn="") => {
	let nip = $decodeBase64(snip)
	let idx = "filename_"+fileIdx;
	
	var fileDisplay = await $fetch(`/api/fileserver/static/dokumen/${nip}/${dataJabatan.value[idx]}`)
	selUploadDoc.value = fileIdx
	selUploadDocBkn.value = idxBkn
	selUploadDocJudul.value = judul
	if(fileDisplay != null){
		const blobUrl = URL.createObjectURL(fileDisplay)
		modalUploadDoc.showModal(`Upload Dokumen ${judul}`, "application/pdf", blobUrl)	
	}else{
		modalUploadDoc.showModal(`Upload Dokumen ${judul}`, "application/pdf", null)	
	}
}

const callback = async(e) => {
	const {name} = e

	if(name == 'doAction'){
		showSpinner.value = true
		let formData = new FormData();
		let nip = $decodeBase64(snip)
		let id = $decodeBase64(sid)

		formData.append("id", id);
		formData.append("file", fileBlob.value[0]);
		formData.append("field_name", "filename_"+selUploadDoc.value)
		formData.append("filename", nip+"_"+selUploadDoc.value+id);
		formData.append("path", 'dokumen/'+nip);

		formData.append("id_ref_dokumen", selUploadDocBkn.value)
		formData.append("idSync", dataJabatan.value.idSync)

		var {data, error} = await useFetch(`/api/uploads/upload/jabatan/${nip}`, {
			method: 'POST',
			body: formData,
		});


		if(data.value.status != "200"){
			const {title, description} = data.value.data
			toast.add({
				id: 'error_post_upload_jabatan',
				icon: "i-heroicons-x-circle",
				title: title,
				description: description,
				timeout: 6000,
				color: 'red'
			}) 	
		}else{
			toast.add({
				id: 'success_post_upload_jabatan',
				icon: "i-heroicons-check-circle",
				title: `Berhasil Upload Dokumen ${selUploadDocJudul.value}`,
				description: `Dokumen berhasil di Upload`,
				timeout: 6000
			}) 	
		}

		refreshNuxtData(["getDataRef", "getDataJabatan"])

		showSpinner.value = false
	}
}
</script>
<template>
	<AppLoadingSpinner :show="showSpinner" />
	<div class="mx-auto">
		<!-- Card -->
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2">
			<div class="flex justify-end mb-4" v-if="!route.query.akhir">
				<button type="button" class="flex justify-center items-center size-7 text-sm font-semibold border border-transparent text-gray-200 hover:text-gray-800 bg-red-500 hover:bg-red-200 disabled:opacity-50 disabled:pointer-events-none" @click="$router.go(-1)">
					<span class="sr-only">Close</span>
					<svg class="flex-shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
				</button>
			</div>	        
			<div class="mb-8">
				<div class="inline-flex justify-between items-center w-full gap-x-2">
					<h2 class="text-xl font-bold text-blue-600">Jabatan {{route.query.akhir ? 'Terakhir' : ''}}</h2>
					<div class="flex gap-x-2 w-[60%]">
						<input id="nip" type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="NIP" :value="dataJabatan.nip">
						<input id="nama" type="text" class="py-2 px-4 pe-11 block w-full border border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-80 disabled:pointer-events-none bg-white" placeholder="Nama">
					</div>
				</div>
			</div>
			<form>
				<!-- Grid -->
				<div class="grid sm:grid-cols-12 gap-2 gap-2.5">
					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<!--<SearchSelect2 ref="refJnsJab" id="jenis_jabatans" :options="jenis_jabatans" keyList="id" namaList="nama" v-model="dataJabatan.jnsjab"/>-->
						<SearchSelect2 ref="refJnsJab" id="jenis_jabatans" :options="jenis_jabatans" keyList="id" namaList="nama" v-model="jnsjab"/>
					</div>

					<div class="sm:col-span-3  sm:col-start-1">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Jenis Mutasi</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<SearchSelect2 ref="refJnsTugasMutasi" id="jenis_mutasi" :options="jenis_tugas_mutasi" keyList="id" namaList="nama" v-model="dataJabatan.jns_tugas_mutasi"/>
					</div>

					<template v-if="dataJabatan.jns_tugas_mutasi == 'MU'">
						<div class="sm:col-span-2 sm:col-start-8">
							<label class="inline-block text-sm text-gray-800 mt-2.5">TMT Mutasi</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-3">
							<input id="tmt_mutasi" type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmt_mutasi">
						</div>
					</template>

					<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>
					
					<div class="sm:col-span-3 sm:col-start-1">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Unit Kerja</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect3 ref="refOpd" id="opds" apiUrl="/api/gets/opd/filter" keyList="id" namaList="nama" v-model="selopd"/>
						<!--<span class="py-2 px-3 block w-full -mt-px -ms-px text-xs italic">Ref. Unor Id BKN : {{dataJabatan.id_unor_bkn}}</span>-->
					</div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Eselon</label>
					</div>
					<!-- End Col -->
					<div class="sm:col-span-3">
						<SearchSelect2 ref="refEselon" id="eselons" :options="eselons" keyList="id" namaList="nama" v-model="dataJabatan.keselon" disabled="true" />
					</div>

					<template v-if="([2 ,4, 5 ,'2','4', '5']).includes(dataJabatan.jnsjab)">
						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<!--<SearchSelect2 ref="refJabatan" id="jabatans" :options="jabatans" keyList="id" namaList="nama" v-model="selkjab"/>-->
							<template v-if="dataJabatan.jnsjab == '2'">
								<template v-if="dataJabatan.jns_tugas_mutasi == 'MJ' ">
									<SearchSelect3 ref="refJabatan" id="jabatans_ft" apiUrl="/api/gets/jabatan_ft_bkn" v-model="selkjab"/>
								</template>
								<template v-else-if="dataJabatan.jns_tugas_mutasi =='MU' ">
									<SearchSelect2 ref="refJabatan" id="riw_jabatans_ft" :options="jabatans" keyList="keyList" namaList="namaList" v-model="selkjab"/>
								</template>
							</template>
							<template v-else-if="dataJabatan.jnsjab == '4'">
								<template v-if="dataJabatan.jns_tugas_mutasi == 'MJ' ">
									<SearchSelect3 ref="refJabatan" id="jabatans_fu" apiUrl="/api/gets/jabatan_fu_bkn/filter" v-model="selkjab"/>	
								</template>
								<template v-if="dataJabatan.jns_tugas_mutasi =='MU' ">
									<SearchSelect2 ref="refJabatan" id="riw_jabatans_fu" :options="jabatans" keyList="keyList" namaList="namaList" v-model="selkjab"/>
								</template>
							</template>
							<template v-else-if="dataJabatan.jnsjab == '5' && dataJabatan.jns_tugas_mutasi =='MJ' ">
								<SearchSelect3 ref="refJabatan" id="jabatans_negara" apiUrl="/api/gets/jabatan_negara" v-model="selkjab"/>
								<span>jabatan_negara</span>
							</template>
						</div>
					</template>

					<template v-if="([2,'2']).includes(dataJabatan.jnsjab)">
						<div class="sm:col-span-3 sm:col-start-1">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Sub Jabatan</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<SearchSelect2 ref="refSubJab" id="subjabs" :options="subjabs" keyList="id" namaList="nama" v-model="selsubjab"/>
						</div>
					</template>

					<div class="sm:col-span-3 sm:col-start-1">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Nama Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<textarea id="njab" class="py-2 px-3 block w-full border border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" rows="3" placeholder="Nama Jabatan" v-model="dataJabatan.njab"></textarea>
						</div>
					</div>
					<!-- End Col -->
					<div class="sm:col-span-12 my-2 first:pt-0 last:pb-0 border-t first:border-transparent border-gray-200"></div>

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">TMT Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<input id="tmtjab" type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tmtjab">
					</div>
					<!-- End Col -->

					<div class="sm:col-span-2 sm:col-start-8">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Tanggal Pelantikan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<input id="tlantik" type="date" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tlantik">
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label class="inline-block text-sm text-gray-800 mt-2.5">Pejabat Penetap</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<SearchSelect2 ref="refPejabat" id="pejabats" :options="pejabats" keyList="kpej" namaList="npej" v-model="dataJabatan.kpej"/>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-3">
						<label for="nskjabat" class="inline-block text-sm text-gray-800 mt-2.5">SK Jabatan</label>
					</div>
					<!-- End Col -->

					<div class="sm:col-span-9">
						<div class="sm:flex gap-2">
							<input id="nskjabat" type="text" name="nskjabat" class="py-2 px-3 block w-full border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="dataJabatan.nskjabat">
							<input id="tskjabat" name="tskjabat" type="date" class="py-2 px-3 block w-[30%] border border-gray-200 shadow-sm -mt-px -ms-px rounded-lg text-sm relative focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none bg-white" v-model="tskjabat">
						</div>
					</div>
					<!-- End Col -->


					<template v-if="!route.query.akhir">
						<div class="sm:col-span-3">
							<label class="inline-block text-sm text-gray-800 mt-2.5">Jabatan Terakhir</label>
						</div>
						<!-- End Col -->

						<div class="sm:col-span-9">
							<div class="sm:flex gap-2">
								<div class="flex mt-2">
										<input id="akhir" type="checkbox" class="shrink-0 mt-0.5 border-gray-200 rounded text-blue-600 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" v-model="akhir">
										<label for="akhir" class="text-sm text-gray-500 ms-3">{{aakhir}}</label>
								</div>
							</div>
						</div>
						<!-- End Col -->
					</template>
				</div>
				<div class="mt-5 grid sm:grid-cols-12 gap-2">
					<div class="sm:col-span-12 sm:col-start-4">
						<div class="sm:flex gap-2">
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" @click="refresh">
									Batal
							</button>
							<template v-if="route.query.akhir">
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-yellow-500 text-white hover:bg-yellow-600 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData('Baru', 0)">
										Jabatan Akhir Baru [0]
								</button>	
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-yellow-600 text-white hover:bg-yellow-700 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData('Baru', 1)">
										Jabatan Akhir Baru [1]
								</button>	
								<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-yellow-700 text-white hover:bg-yellow-800 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData('Baru', 2)">
										Jabatan Akhir Baru [2]
								</button>	
							</template>
							<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-500 text-white hover:bg-blue-600 disabled:opacity-50 disabled:pointer-events-none" @click="simpanData(method)">
								{{method}} Data
							</button>	
						</div>					
					</div>
				</div>
			</form>
		</div>
		<!-- End Card -->
		<div class="bg-white rounded-xl shadow py-4 px-6 border-t-2 mt-6" v-if="method == 'Update'">
			<div class="mb-8">
				<h2 class="text-xl font-bold text-blue-600">Dokumen Jabatan</h2>
			</div>
			<div class="grid sm:grid-cols-4 gap-2 gap-2.5">
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('SK Jabatan','jab', '872')">
					SK Jabatan
				</button>
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('Surat Pelatikan','lantik', '873')">
					Surat Pelantikan
				</button>
				<button type="button" class="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" @click="doUploadDoc('SK Penugasan','tugas_mutasi')">
					SK {{dataJabatan.jnsjab == 1 ? 'Penugasan' : 'Mutasi'}}
				</button>
			</div>
		</div>
	</div>
	<!-- End Card Section -->
</template>