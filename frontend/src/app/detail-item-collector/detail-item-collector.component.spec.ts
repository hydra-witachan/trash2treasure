import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailItemCollectorComponent } from './detail-item-collector.component';

describe('DetailItemCollectorComponent', () => {
  let component: DetailItemCollectorComponent;
  let fixture: ComponentFixture<DetailItemCollectorComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DetailItemCollectorComponent]
    });
    fixture = TestBed.createComponent(DetailItemCollectorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
