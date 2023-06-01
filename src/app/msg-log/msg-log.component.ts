import { FormBuilder } from '@angular/forms';

import { RequestService } from '../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from '../base/table';
import { FormGroup, Validators } from '@angular/forms';
@Component({
    selector: 'app-msg-log',
    templateUrl: './msg-log.component.html',
    styleUrls: ['./msg-log.component.scss'],
})
export class MsgLogComponent {
    constructor(
        private router: Router,
        private rs: RequestService,
        private msg: NzMessageService,
        private fb: FormBuilder
    ) {
        this.load();
        this.build();
    }

    build(obj?: any) {
        obj = obj || {};
        this.group = this.fb.group({
            // id: [obj.id || '', []],
            cellphone: [obj.cellphone || '', []],
            content: [obj.content || '', []],
            result: [obj.result || false, []],
            type: [obj.type || '', []],
        });
    }
    group!: FormGroup;
    listOfOption = [
        { value: 'aliyun', label: '阿里云' },
        { value: 'tccloud', label: '腾讯云' },
        { value: 'hicloud', label: '华为云' },
    ];
    url = 'sms/api/sms/';
    isVisible = false;
    addVisible = false;
    loading = true;
    datum: any[] = [];
    total = 1;
    pageSize = 20;
    pageIndex = 1;
    query: any = {};
    load() {
        this.loading = true;
        this.rs
            .post(this.url + 'search', this.query)
            .subscribe((res) => {
                this.datum = res.data;
                this.total = res.total;
            })
            .add(() => {
                this.loading = false;
            });
    }
    delete( id: number) {
        this.rs.get(this.url + `${id}/delete`).subscribe((res) => {
            this.msg.success('删除成功');
            this.load();
        });
    }
    onAdd() {
        this.isVisible = true;
    }

    onQuery($event: NzTableQueryParams) {
        ParseTableQuery($event, this.query);
        this.load();
    }
    pageIndexChange(pageIndex: number) {
        this.query.skip = pageIndex - 1;
    }
    pageSizeChange(pageSize: number) {
        this.query.limit = pageSize;
    }
    search(text: any) {
        if (text)
            this.query.filter = {
                id: text,
            };
        else this.query = {};
        this.load();
    }
    cancel() {
        //  this.msg.info('取消删除');
        this.isVisible = false;
    }
   disDelete(){this.msg.info('取消删除');}
    open(id: string) {
        //this.router.navigateByUrl("/admin/device/" + id)
    }
    
    submit() {
        if (this.group.valid) {
            console.log(this.group.value);
            // let url = this.id ? `device/type/${this.id}` : `device/type/create`;
            this.rs
                .post(this.url + 'create', this.group.value)
                .subscribe((res) => {
                    this.msg.success('保存成功');
                    this.isVisible = false;
                    this.load();
                });

            return;
        } else {
            Object.values(this.group.controls).forEach((control) => {
                if (control.invalid) {
                    control.markAsDirty();
                    control.updateValueAndValidity({ onlySelf: true });
                }
            });
        }
    }
}
