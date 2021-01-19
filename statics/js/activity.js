var tableHeight = document.body.clientHeight - 120;
var date = formatDate(new Date());
new Vue({
    el: '#app',
    data: {
        cosClient:null,
        todayStr:date,
        height:tableHeight,
        total:0,
        page_size:10,
        current_page:1,
        createActivity:true,
        giftList:[],
        gifts:[],
        activity:{
            name:"",
            type:3,
            gift_id:0,
            limit_join:0,
            join_limit_num:1,
            receive_limit:1,
            des:"",
            attachments:[],
            start_at:"",
            end_at:"",
            run_at:"",
            share_title:"",
            share_image:"",
        },
        cos:{
            token:"",
            domain:"",
            bucket:"",
            region:"",
            tmp_secret_id:"",
            tmp_secret_key:"",
            start_time:"",
            expired_time:"",
            env:""
        },
        rules: {
            name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
            num: [{ required: true, message: '请输入数量', trigger: 'blur' }],
            type: [{ required: true, message: '请输入数量', trigger: 'blur' }],
            status: [{ required: true, message: '请输入数量', trigger: 'blur' }],
            attachments: [{ required: true, message: '请输入数量', trigger: '' }],
        },
    },
    created: function () {
        this.getPage();
        this.getToken();
        this.getGifts();
    },
    methods: {
        getGifts:function (e) {
            var url = "/admin/api/gift/enable_list";
            axios.get(url,{}).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    let data = [];
                    this.gifts = res.data;
                    if (this.gifts.length > 0){
                        this.activity.gift_id = this.gifts[0]["ID"]
                    }
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        getPage:function (e) {
            var url = "/admin/api/gift/page";
            axios.get(url+"?page_size="+this.page_size+'&page_num='+this.current_page+'&order_by=created_at&sort=desc',{
                page_size:this.page_size,
                page_number:this.current_page,
                order_by:'created_at',
                sort_by:'desc'
            }).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    let data = [];
                    this.giftList = res.data
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        deleteRow(index, rows) {
            rows.splice(index, 1);
        },
        cancelCreate(done) {
            this.$confirm('确认关闭？')
                .then(_ => {
                    done();
                })
                .catch(_ => {});
        },
        //提交数据
        submitActivity(){
            var url = "/admin/api/activity/create";
            this.activity.join_limit_num = parseFloat(this.activity.join_limit_num);
            axios.post(url,this.activity).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.$message.success(res.msg);
                    this.createGift = false
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        //上传图片到cos
        upload(e){
            var timestamp=new Date().getTime();
            let fileName = this.cos.env+"/admin/"+this.todayStr+"/"+timestamp+"_"+e.file.name;
            this.cosClient.putObject({
                Bucket: this.cos.bucket,        /* 必须 */
                Region: this.cos.region,        /* 存储桶所在地域，必须字段 */
                Key: fileName,                  /* 必须 */
                StorageClass: 'STANDARD',
                Body: e.file,                   // 上传文件对象
                onProgress: progressData=> {
                    console.log(JSON.stringify(progressData));
                }
            }, (err, data)=> {
                if (err == null){
                    if (e.data == 1){
                        this.activity.share_image = [fileName]
                    }else{
                        this.activity.attachments = [fileName]
                    }
                }else{
                    this.$message.error("图片上传出错");
                }
            });
        },
        //处理上传前
        uploadPreview(e){

        },
        //处理上传后
        uploadRemove(e){
            if (e.data == 1){
                this.activity.share_image = []
            }else{
                this.activity.attachments = []
            }
        },
        getToken:function (e) {
            var url = "/admin/api/cos/token";
            axios.get(url).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.cos.token          = res.data.token;
                    this.cos.domain         = res.data.domain;
                    this.cos.bucket         = res.data.bucket;
                    this.cos.region         = res.data.region;
                    this.cos.tmp_secret_id  = res.data.tmp_secret_id;
                    this.cos.tmp_secret_key = res.data.tmp_secret_key;
                    this.cos.start_time     = res.data.start_time;
                    this.cos.expired_time   = res.data.expired_ime;
                    this.cos.env            = res.data.env;
                    let cos = this.cos
                    this.cosClient = new COS({
                        // 必选参数
                        getAuthorization: function (options, callback) {
                            callback({
                                TmpSecretId: cos.tmp_secret_id,
                                TmpSecretKey: cos.tmp_secret_key,
                                XCosSecurityToken: cos.token,
                                StartTime: cos.start_time,
                                ExpiredTime: cos.expired_time,
                            });
                        }
                    });
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        }
    }
})