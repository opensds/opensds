import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { Validators, FormControl, FormGroup, FormBuilder } from '@angular/forms';

import { Message, SelectItem } from './../../../components/common/api';

import { VolumeService ,ReplicationService} from './../volume.service';
import { ProfileService } from './../../profile/profile.service';

@Component({
  selector: 'app-create-volume',
  templateUrl: './create-volume.component.html',
  styleUrls: [

  ],
  animations: [
    trigger('overlayState', [
      state('hidden', style({
        opacity: 0
      })),
      state('visible', style({
        opacity: 1
      })),
      transition('visible => hidden', animate('400ms ease-in')),
      transition('hidden => visible', animate('400ms ease-out'))
    ]),

    trigger('notificationTopbar', [
      state('hidden', style({
        height: '0',
        opacity: 0
      })),
      state('visible', style({
        height: '*',
        opacity: 1
      })),
      transition('visible => hidden', animate('400ms ease-in')),
      transition('hidden => visible', animate('400ms ease-out'))
    ])
  ]
})
export class CreateVolumeComponent implements OnInit {

  bbbbb = 'name';
  label = {};
  availabilityZones = [];
  volumeform;
  volumeItems = [0];
  capacityUnit = [];
  profileOptions = [];
  capacity = 'GB';
  createVolumes = [];
  value: boolean;
  showReplicationConf = false;

  errorMessage = {
    "zone": { required: "Zone is required."}
  };
    defaultProfile = {
        label: null,
        value: {id:null,profileName:null}
    };
  constructor(
    private router: Router,
    private fb: FormBuilder,
    private ProfileService: ProfileService,
    private VolumeService: VolumeService,
    private replicationService:ReplicationService
  ) {}

  ngOnInit() {
    this.label = {
      zone: 'Availability Zone',
      name: 'Name',
      profile: 'Profile',
      capacity: 'Capacity',
      quantity: 'Quantity'
    }

    this.availabilityZones = [
      {
        label: 'Default', value: 'default'
      }
    ];

    this.getProfiles();

    this.capacityUnit = [
      {
        label: 'GB', value: 'GB'
      },
      {
        label: 'TB', value: 'TB'
      }
    ];
    this.volumeform = this.fb.group({
      'zone': new FormControl('default', Validators.required),
      'name0': new FormControl('', Validators.required),
      'profileId0': new FormControl(this.defaultProfile, Validators.required),
      'size0': new FormControl(1, Validators.required),
      'capacity0': new FormControl(''),
      'quantity0': new FormControl(1)
    });
    this.volumeform.valueChanges.subscribe(
      (value:string)=>{
          this.createVolumes = this.getVolumesDataArray(this.volumeform.value);
          this.setRepForm();
      }
    );
      this.createVolumes = this.getVolumesDataArray(this.volumeform.value);
      this.setRepForm();
  }

  addVolumeItem() {
    this.volumeItems.push(
      this.volumeItems[this.volumeItems.length-1] + 1
    );
    this.volumeItems.forEach(index => {
      if(index !== 0){
        this.volumeform.addControl('name'+index, this.fb.control('', Validators.required));
        this.volumeform.addControl('profileId'+index, this.fb.control(this.defaultProfile,Validators.required));
        this.volumeform.addControl('size'+index, this.fb.control(1, Validators.required));
        this.volumeform.addControl('capacity'+index, this.fb.control('GB', Validators.required));
        this.volumeform.addControl('quantity'+index, this.fb.control(1));
      }
    });
  }

  getProfiles() {
    this.ProfileService.getProfiles().subscribe((res) => {
      let profiles = res.json();
      profiles.forEach(profile => {
        this.profileOptions.push({
          label: profile.name,
          value: {id:profile.id,profileName:profile.name}
        });
      });
    });
  }

  deleteVolumeItem(index) {
      this.volumeItems.splice(index, 1);
      this.volumeform.removeControl('name'+index);
      this.volumeform.removeControl('profileId'+index);
      this.volumeform.removeControl('size'+index);
      this.volumeform.removeControl('capacity'+index);
      this.volumeform.removeControl('quantity'+index);
  }

  createVolume(param){
    this.VolumeService.createVolume(param).subscribe((res) => {
      this.router.navigate(['/block']);
    });
  }
  createVolumeAndReplication(volParam,repParam){
    this.VolumeService.createVolume(volParam).subscribe((res2) => {
        this.VolumeService.createVolume(repParam).subscribe((res) => {
            let param = {
                "name":res.json().name ,
                "primaryVolumeId": res2.json().id,
                "availabilityZone": res.json().availabilityZone,
                "profileId": res.json().profileId,
                "replicationMode":"async",
                "replicationPeriod":this.createVolumes["formGroup"].value.period,
                "secondaryVolumeId":res.json().id
            }
            this.replicationService.createReplication(param).subscribe((res) => {});
            this.router.navigate(['/block']);
        });
    });
  }
  onSubmit(value) {
      if(!this.volumeform.valid){
          for(let i in this.volumeform.controls){
              this.volumeform.controls[i].markAsTouched();
          }
          return;
      }
      if(this.showReplicationConf && !this.createVolumes["formGroup"].valid){
          for(let i in this.createVolumes["formGroup"].controls){
              this.createVolumes["formGroup"].controls[i].markAsTouched();
          }
          return;
      }
      let dataArr = this.getVolumesDataArray(value);
      let volumeData = [];
      dataArr.forEach(item => {
          volumeData.push({
              name: item.name,
              size: item.size,
              availabilityZone: item.availabilityZone,
              profileId: item.profile.id
          });
      });
      for(let i in volumeData){
          if(this.showReplicationConf){
              let repVolume = {
                  name:null,
                  profileId:null
              };
              Object.assign(repVolume,volumeData[i]);
              repVolume.name = this.createVolumes["formGroup"].value["name"+i];
              repVolume.profileId = this.createVolumes["formGroup"].value["profileId"+i];
              this.createVolumeAndReplication(volumeData[i],repVolume);
          }else{
              this.createVolume(volumeData[i]);
          }
      }
  }
  getVolumesDataArray(value){
      let dataArr = [];
      this.volumeItems.forEach(index => {
          if(!value['capacity'+index]){
              value['capacity'+index]='GB';
          }
          let unit = value['capacity'+index]==='GB' ? 1 : 1024;
          let qunantity = value['quantity'+index];
          if(qunantity && qunantity !== 1){
              for(let i=0;i<qunantity;i++){
                  dataArr.push({
                      name: value['name'+index]+i,
                      size: value['size'+index]*unit,
                      availabilityZone: value.zone,
                      profile: value['profileId'+index]
                  });
              }
          }else{
              dataArr.push({
                  name: value['name'+index],
                  size: value['size'+index]*unit,
                  availabilityZone: value.zone,
                  profile: value['profileId'+index]
              });
          }
      });
      return dataArr;
  }
    checkRep(param:boolean){}
    //create replication volumes formGroup
    setRepForm(){
        let param = {
            'zone': new FormControl(this.createVolumes[0].availabilityZone, Validators.required),
            'period': new FormControl(60, Validators.required)
        };
        for(let i in this.createVolumes){
            param["name"+i] = new FormControl(this.createVolumes[i].name+"-replication", Validators.required);
            param["profileId"+i] = new FormControl('', Validators.required);
        }
        this.createVolumes["formGroup"] = this.fb.group(param);
    }

}
