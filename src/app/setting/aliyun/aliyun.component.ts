import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from 'src/app/request.service';
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
    selector: 'app-aliyun',
    templateUrl: './aliyun.component.html',
    styleUrls: ['./aliyun.component.scss']
})
export class AliyunComponent implements OnInit {
    group!: FormGroup;
    loading = false
    query = {}
    dbData = []

    constructor(private fb: FormBuilder,
                private router: Router,
                private route: ActivatedRoute,
                private rs: RequestService,
                private msg: NzMessageService) {
    }

    switchValue = false;

    ngOnInit(): void {
        this.rs.get(`config`).subscribe(res => {
            //let data = res.data;
            this.build(res.data)
        })

        this.build()
    }


    build(obj?: any) {
        obj = obj || {}
        this.group = this.fb.group({
            id: [obj.id || '', []],
            secret: [obj.secret || ''],
            sign: [obj.sign || ''],
            template: [obj.template || ''],
        })
    }

    submit() {
        if (this.group.valid) {

            this.rs.post(`config`, this.group.value).subscribe(res => {
                this.msg.success("保存成功")
            })

            return;
        } else {
            Object.values(this.group.controls).forEach(control => {
                if (control.invalid) {
                    control.markAsDirty();
                    control.updateValueAndValidity({onlySelf: true});
                }
            });

        }
    }


}

