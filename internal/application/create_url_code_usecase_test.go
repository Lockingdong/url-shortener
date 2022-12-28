package application

import (
	"UrlShortener/internal/adapter/output_port"
	"UrlShortener/internal/domain/repository"
	"context"
	"reflect"
	"testing"
)

func TestCreateUrlCodeUseCase_Execute(t *testing.T) {
	type fields struct {
		createUrlInfoRepository repository.ICreateUrlCodeRepository
	}
	type args struct {
		ctx context.Context
		cmd *CreateUrlCodeCommand
	}
	tests := []struct {
		args    args
		fields  fields
		want    *CreateUrlCodeResponse
		name    string
		wantErr bool
	}{
		{
			name: "TestCreatUrlCode",
			fields: fields{
				createUrlInfoRepository: output_port.NewUrlInfoMockRepository(nil),
			},
			args: args{
				ctx: context.Background(),
				cmd: &CreateUrlCodeCommand{
					Url: "https://dongstudio.medium.com/",
				},
			},
			want: &CreateUrlCodeResponse{
				UrlCode: "test_code",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &CreateUrlCodeUseCase{
				repository: tt.fields.createUrlInfoRepository,
			}
			got, err := u.Execute(tt.args.ctx, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUrlCodeUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if reflect.TypeOf(got.UrlCode) != reflect.TypeOf(tt.want.UrlCode) {
				t.Errorf("CreateUrlCodeUseCase.Execute() = %v, want %v", reflect.TypeOf(got.UrlCode), reflect.TypeOf(tt.want.UrlCode))
			}
		})
	}
}
