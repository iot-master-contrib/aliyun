 
 
  import { Component, OnInit } from '@angular/core';
  import { FormBuilder, FormGroup, Validators } from "@angular/forms";
  import { ActivatedRoute, Router } from "@angular/router";
  import { RequestService } from 'src/app/request.service';
  import { NzMessageService } from "ng-zorro-antd/message";
  

  @Component({
    selector: 'app-tc-cloud',
    templateUrl: './tc-cloud.component.html',
    styleUrls: ['./tc-cloud.component.scss']
  })
  export class TcCloudComponent implements OnInit {
    group!: FormGroup;
    loading=false
    query={}
    dbData=[] 
    constructor(private fb: FormBuilder,
      private router: Router,
      private route: ActivatedRoute,
     private rs: RequestService,
      private msg: NzMessageService) {this.load()
    }
    switchValue = false;
  
    ngOnInit(): void {
      this.rs.get(`config`).subscribe(res => {
        //let data = res.data;
        this.build(res.data)
      })
  
      this.build()
    }
  
    load(){ 
      this.rs.get(`config/database`).subscribe((res) => {   
       this.dbData=res.data
      // this.group.patchValue({  })
      }); 
  }
  
    build(obj?: any) {
      obj = obj || {}
      this.group = this.fb.group({
        AccessKeyId: [obj.AccessKeyId || '', []],
        AccessKeySecret: [obj.AccessKeySecret||false ] 
      })
    }
  
    submit() {
      if (this.group.valid) {
         
        this.rs.post(`config/database`, this.group.value).subscribe(res => {
          this.msg.success("保存成功")
        })
  
        return;
      }
      else {
        Object.values(this.group.controls).forEach(control => {
          if (control.invalid) {
            control.markAsDirty();
            control.updateValueAndValidity({ onlySelf: true });
          }
        });
  
      }
    }
     
     
  }
  
