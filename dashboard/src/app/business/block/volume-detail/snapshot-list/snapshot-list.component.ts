import { Component, OnInit, Input } from '@angular/core';
import { FormControl, FormGroup, FormBuilder, Validators, ValidatorFn, AbstractControl } from '@angular/forms';
import { VolumeService,SnapshotService } from './../../volume.service';
import { ConfirmationService,ConfirmDialogModule} from '../../../../components/common/api';

@Component({
  selector: 'app-snapshot-list',
  templateUrl: './snapshot-list.component.html',
  styleUrls: [

  ]
})
export class SnapshotListComponent implements OnInit {

  @Input() volumeId;
  volume;
  label;
  selectedSnapshotId;
  selectedSnapshots = [];
  snapshortfilter;
  snapshots;
  snapshotfilter;
  snapshotPropertyDisplay = false;
  snapshotFormGroup;

  isCreate = false;
  isModify = false;
  snapshotProperty = {
    name: '',
    description: ''
  }
  okBtnDisabled = false;

  errorMessage = {
      "name": { required: "Name is required." },
      "description": { maxlength: "Max. length is 200." }
  };

  constructor(
    private VolumeService: VolumeService,
    private SnapshotService: SnapshotService,
    private fb: FormBuilder,
    private confirmationService:ConfirmationService
  ) {
    this.snapshotFormGroup = this.fb.group({
      "name": ["", Validators.required],
      "description": ["", Validators.maxLength(200)]
    });
  }

  ngOnInit() {
    this.getVolumeById(this.volumeId);
    this.label = {
      name: 'name',
      volume: 'Volume',
      description: 'Description'
    }
    this.getSnapshots(
      {
        key: 'volumeId',
        value: this.volumeId
      }
    );
  }

  getVolumeById(volumeId){
    this.VolumeService.getVolumeById(volumeId).subscribe((res) => {
      this.volume = res.json();
    });
  }

  createSnapshot() {
    let param = {
      name: this.snapshotFormGroup.value.name,
      volumeId: this.volumeId,
      description: this.snapshotFormGroup.value.description
    }
    this.SnapshotService.createSnapshot(param).subscribe((res) => {
      this.getSnapshots(
        {
          key: 'volumeId',
          value: this.volumeId
        }
      );
    });
  }

  batchDeleteSnapshot(param) {
    if (param) {
        let  msg;
        if(param.length>1){

            msg = "<div>Are you sure you want to delete the selected Snapshots?</div><h3>[ "+ param.length +" Snapshots ]</h3>";
        }else{
            msg = "<div>Are you sure you want to delete the Snapshot?</div><h3>[ "+ param[0].name +" ]</h3>";
        }

        this.confirmationService.confirm({
            message: msg,
            header: "Delete Snapshots",
            acceptLabel: "Delete",
            isWarning: true,
            accept: ()=>{
                param.forEach(snapshot => {
                    this.deleteSnapshot(snapshot.id);
                });

            },
            reject:()=>{}
        })
    }
  }

  deleteSnapshot(id) {
    this.SnapshotService.deleteSnapshot(id).subscribe((res) => {
      this.getSnapshots(
        {
          key: 'volumeId',
          value: this.volumeId
        }
      );
    });
  }

  getSnapshots(filter?) {
    this.SnapshotService.getSnapshots(filter).subscribe((res) => {
      this.snapshots = res.json();
      this.snapshotPropertyDisplay = false;
    });
  }

  modifySnapshot(){
    let param = {
      name: this.snapshotFormGroup.value.name,
      description: this.snapshotFormGroup.value.description
    }
    this.SnapshotService.modifySnapshot(this.selectedSnapshotId,param).subscribe((res) => {
      this.getSnapshots(
        {
          key: 'volumeId',
          value: this.volumeId
        }
      );
    });
  }

  showSnapshotPropertyDialog(method,selectedSnapshot?){
    this.snapshotPropertyDisplay = true;
    if(method === 'create'){
      this.isCreate = true;
      this.isModify = false;
      this.snapshotProperty.name = '';
      this.snapshotProperty.description = '';
    }else if(method === 'modify'){
      this.isCreate = false;
      this.isModify = true;
      this.snapshotProperty.name = selectedSnapshot.name;
      this.snapshotProperty.description = selectedSnapshot.description;
    }
    if(selectedSnapshot && selectedSnapshot.id){
      this.selectedSnapshotId = selectedSnapshot.id;
    }
  }

  snapshotModifyOrCreate(){
    if(this.isModify){
      this.modifySnapshot();
    }else{
      this.createSnapshot();
    }

  }

}
