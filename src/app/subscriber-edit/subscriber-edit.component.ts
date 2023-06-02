import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from 'src/app/request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
    selector: 'app-subscriber-edit',
    templateUrl: './subscriber-edit.component.html',
    styleUrls: ['./subscriber-edit.component.scss'],
})
export class SubscriberEditComponent implements OnInit {
    group!: FormGroup;
    id: any = 0;
    add: string = 's';
    @ViewChild('parametersChild') parametersChild: any;
    listOfOption: any[] = [
        { value: 'aa', label: 11 },
        { value: 'bb', label: 12 },
    ];
    productList: any[] = [
        { value: 'cc', label: 21 },
        { value: 'dd', label: 22 },
    ];

    constructor(
        private fb: FormBuilder,
        private router: Router,
        private route: ActivatedRoute,
        private rs: RequestService,
        private msg: NzMessageService
    ) {}

    ngOnInit(): void {
        if (this.route.snapshot.paramMap.has('id')) {
            this.id = this.route.snapshot.paramMap.get('id');
            this.rs.get(`subscriber/${this.id}`).subscribe((res) => {
                //let data = res.data;
                this.build(res.data);
            });
        }

        this.build();

        this.rs
            .post('subscriber/search', {})
            .subscribe((res) => {
                const data: any[] = [];

                res.data.filter((item: { id: string; name: string }) =>
                    data.push({
                        label: item.id + ' / ' + item.name,
                        value: item.id,
                    })
                );
                this.productList = data;
            })
            .add(() => {});
    }

    build(obj?: any) {
        obj = obj || {};
        this.group = this.fb.group({
            id: [obj.id || '', []],
            name: [obj.name || '', []],
            cellphone: [obj.cellphone || '', []],
            level: [obj.level || 1, []],
            disabled: [obj.disabled || false, []],
        });
    }


      parameterAdd($event:any){

        $event.stopPropagation()
    if (this.parametersChild) {
      this.parametersChild.group.get('keyName').controls.unshift(
        this.fb.group({
            assign: ['', []],
            expression: ['', []],
         type: ['first', []]
        })
      )
    }

      }
      getValueData() {

        const parametersGroup = this.parametersChild.group;
        const parameters = parametersGroup.get('keyName').controls.map((item: { value: any; }) => item.value);
        return    parameters   ;

      }
    submit() {
        let value = this.group.value;
        if (this.group.valid) {
            let url = this.id
                ? `subscriber/${this.id}`
                : `subscriber/create`;

           const sendData = JSON.parse(JSON.stringify(this.group.value));
           sendData.aggregators = this.getValueData();

            this.rs.post(url, sendData).subscribe((res) => {
                this.handleCancel();
                this.msg.success('保存成功');
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
    handleCancel() {
        const path = `/subscriber`;
        this.router.navigateByUrl(path);
    }
}
