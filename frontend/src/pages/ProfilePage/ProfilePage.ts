import Base from '@components/Base/Base';
import Header from '@components/Header/Header';
import { validate } from '@utils/validation';
import AuthorizationForm from '@components/Form/AuthorizationForm';
import JourneyPreview from './JourneyPreview';
import { authorize, editProfile, getUserProfile, imageUpload, resetPassword } from '@api/user';
import { profileErrors, signupErrors, loginErrors } from '../../types/errors';
import { ROUTES } from '@router/ROUTES';
import ProfileBlock from './ProfileBlock';
import { router } from '@router/router';
import template from '@templates/ProfilePage.hbs';
import { getUserTrips } from '@api/journey';
import { getUserAlbumsByUserID } from '@api/album';
import { UserAuthResponseData } from '@types/api';
import AlbumPreview from './AlbumPreview';
import { AlbumData } from '@types/api';

class ProfilePage extends Base {

  isOwn: boolean;

  userID: number;

  form: AuthorizationForm;

  constructor(parent: HTMLElement, params : unknown) {
    super(parent, template);
    this.userID = parseInt(params[0]);
  }

  emptyContent(entityName : string, parent : HTMLDivElement) {
    this.createElement('h3', {}, this.isOwn ? `Вы пока не создавали ${entityName}` : `Пользователь пока не создавал ${entityName}`, {
      parent: parent, position: 'into',
    });
  }

  render() {

    if (!this.userID) {
      router.go('404');
      return;
    }

    getUserProfile(this.userID).then((profileData) => {

      this.isOwn = (this.userData === null) ? false : (this.userData.userID === profileData.data.id);
      this.preRender();

      if (profileData.data.id === 0) {
        router.go('404');
        return;
      }

      const authForm = new AuthorizationForm(this.parent, '');

      const profileTemplateData = {
        userID: profileData.data.id, username: profileData.data.username, status: profileData.data.bio, avatar: profileData.data.avatar,
      };


      const header = document.getElementById('header') as HTMLElement;
      document.body.classList.remove('auth-background');

      new Header(header).render();

      const profileBlock = document.querySelector('#profile-block') as HTMLDivElement;

      new ProfileBlock(profileBlock, profileTemplateData).render();

      const profileContent = document.querySelector('.profile-content') as HTMLDivElement;

      const profileEditForm = document.querySelector('dialog') as HTMLDialogElement;

      const submitButton = document.getElementById('button-submit') as HTMLButtonElement;

      const cancelButton = document.querySelector('#button-cancel') as HTMLButtonElement;
      cancelButton.addEventListener('click', () => profileEditForm.close());
    
      const passwordInputs = document.querySelectorAll('.password') as NodeListOf<HTMLInputElement>;


      const linkBlock = document.querySelector('#underlined-links') as HTMLDivElement;
      const journeyLink = this.createElement('label', {
        class: 'underlined-link active',
      }, 'Поездки', {
        parent: linkBlock, position: 'into',
      });
      // const albumsLink = this.createElement('label', {
      //   class: 'underlined-link',
      // }, 'Альбомы', {
      //   parent: linkBlock, position: 'into',
      // });

      let JOURNEY_DATA;
      let ALBUM_DATA;

      const contentBlock = document.getElementById('content-block') as HTMLDivElement;
      let createButton: HTMLButtonElement;

      if (this.isOwn) {
        createButton = this.createElement('button', {
          class: 'button-primary button-link', id: 'create-button', href: '/journey/new',
        }, 'Создать поездку', {
          parent: contentBlock, position: 'after',
        }) as HTMLButtonElement;
      }

      journeyLink.addEventListener('click', () => {
        profileContent.innerHTML = '';
        if (!JOURNEY_DATA) {
          this.createElement('h3', {}, this.isOwn ? 'Вы пока не создавали поездки' : 'Пользователь пока не создавал поездки', {
            parent: profileContent, position: 'into',
          });
        } else {
          JOURNEY_DATA.forEach((journey) => new JourneyPreview(profileContent, journey).render());
        }

        
        // albumsLink.classList.remove('active');
        journeyLink.classList.add('active');
        if (createButton) {
          createButton.textContent = 'Создать новую поездку';
          createButton.setAttribute('href', '/journey/new');
        }
      });


      // albumsLink.addEventListener('click', () => {
      //   profileContent.innerHTML = '';

      //   if (ALBUM_DATA && ALBUM_DATA.length > 0) {
      //     ALBUM_DATA.forEach((album : AlbumData) => {
      //       new AlbumPreview(profileContent, album).render();

      //     } );
      //   } else {
      //     this.emptyContent('альбомы', profileContent);
      //   }

      //   albumsLink.classList.add('active');
      //   journeyLink.classList.remove('active');

      //   if (createButton) {
      //     createButton.textContent = 'Создать альбом';
      //     createButton.setAttribute('href', '/albums/new');
      //   }
      // });

      // endblock
      passwordInputs.forEach((input: HTMLInputElement) => input.addEventListener('input', () => {
        const parent = input.parentElement as HTMLElement;
        if (input.value.length > 0) {
          validate(input.value, input.type)
            .catch((error) => { authForm.renderError(parent, error.message); });
        }
        authForm.clearError(parent);
      },
      ));
      const inputs = document.querySelectorAll('input') as NodeListOf<HTMLInputElement>;

      const usernameField = inputs[0] as HTMLInputElement;

      usernameField.addEventListener('change', () => {
        authForm.clearError(usernameField.parentElement as HTMLElement);
      });

      var flagCloseEditForm = false;
      const passwordField = inputs[1] as HTMLInputElement;
      const newPasswordField = inputs[2] as HTMLInputElement;
      const lowestInputDiv = passwordField.parentElement as HTMLDivElement;
      const statusField = document.querySelector('textarea') as HTMLTextAreaElement;

      usernameField.value = profileTemplateData.username;
      statusField.value = profileTemplateData.status;

      const imageInput = document.querySelector('#profile-edit-avatar') as HTMLInputElement;
      let formData: FormData;

      let lowestInput = document.querySelectorAll('.input')[4] as HTMLInputElement;

      imageInput.addEventListener('change', () => {
        if (imageInput.files && imageInput.files.length > 0) {
          const file = imageInput.files[0];
          const fileType = file.type;
          const validImageTypes = ['image/gif', 'image/jpeg', 'image/png', 'image/webp'];

          if (!validImageTypes.includes(fileType)) {
            authForm.renderError(lowestInput, 'Выберите аватарку допустимых форматов: jpeg, png, webp, gif');
            return;
          } else {
            authForm.clearError(lowestInput);
          }
          formData = new FormData();
          formData.append('file', file);
          flagCloseEditForm = true;
        }
      });
      // ----------------------------------------------------------------------
      submitButton.addEventListener('click', async (e: Event) => {
        e.preventDefault();

        if (document.querySelectorAll('.has-error').length !== 0) {
          return;
        }

        let error = false;

        try {
          if (formData) {
            const imageUploadResponse = await imageUpload(
              ROUTES.profile.upload(this.userID),
              formData,
            );
            if (imageUploadResponse.status !== 200) {
              authForm.renderError(
                usernameField,
                signupErrors[imageUploadResponse.data.error],
              );
              error = true;
            }
          }

          if (
            passwordField.value === newPasswordField.value &&
            newPasswordField.value.length > 0
          ) {
            authForm.renderError(
              newPasswordField.parentElement as HTMLElement,
              profileErrors.notnew,
            );
            error = true;
          } else if (newPasswordField.value.length > 0) {
            const passwordResponse = await resetPassword(
              this.userData.userID,
              passwordField.value,
              newPasswordField.value,
            );
            if (passwordResponse.status !== 200) {
              authForm.renderError(
                passwordField.parentElement as HTMLElement,
                signupErrors['incorrect old password'],
              );
              error = true;
            } else {
              passwordField.value = '';
              newPasswordField.value = '';
            }
          }

          if (!error) {
            if (usernameField.value.length > 5 || usernameField.value.length === 0) {
              const profileBioNickEditResponse = await editProfile(
                this.userData.userID,
                usernameField.value,
                statusField.value,
              );
              if (profileBioNickEditResponse.status === 200) {
                profileBlock.innerHTML = '';
                const templateData = {
                  userID: profileBioNickEditResponse.data.id,
                  username: profileBioNickEditResponse.data.username,
                  status: profileBioNickEditResponse.data.bio,
                  avatar: !profileBioNickEditResponse.data.avatar ? '' : profileBioNickEditResponse.data.avatar,
                };
                new ProfileBlock(profileBlock, templateData).render();
              } else {
                authForm.renderError(
                  usernameField.parentElement as HTMLElement,
                  profileErrors[profileBioNickEditResponse.data.error],
                );
                error = true;
              }
            } else {
              authForm.renderError(
                usernameField.parentElement as HTMLElement,
                profileErrors.short,
              );
              error = true;
            }
          }

          if (!error) {
            profileEditForm.close();
          }
        } catch (err) {
          console.error('Unexpected error:', err);
          // Дополнительная обработка ошибки, если нужно
          error = true;
        }
      });
      // --------------------------------------------------------------------
      getUserTrips(this.userID).then((journeyList) => {
        if (journeyList.status === 200 && journeyList.data.journeys !== null) {
          JOURNEY_DATA = journeyList.data.journeys;
          profileContent.innerHTML = '';
          JOURNEY_DATA.forEach((journey) => new JourneyPreview(profileContent, journey).render());
          journeyLink.classList.add('active');
        } else {
          this.emptyContent('поездки', profileContent);
        }
      });

      getUserAlbumsByUserID(this.userID).then((albumList) => {
        if (albumList.status === 200 && albumList.data.albums !== null) {
          ALBUM_DATA = albumList.data.albums;
        }
      });

    });

  }
}

export default ProfilePage;
