package builders

import "bot/discord/messages"

type InteractionCallbackBuilder messages.InteractionCallback
type EmbedBuilder messages.Embed
type OptionBuilder messages.Option
type ComponentBuilder messages.Component
type ActionRowComponentBuilder messages.Components
type FieldBuilder messages.Field
type AttachmentBuilder messages.Attachment

func CreateInteractionCallback() *InteractionCallbackBuilder {

	return &InteractionCallbackBuilder{
		Type: 4,
	}
}

func CreateInteractionCallbackEdit(edit bool) *InteractionCallbackBuilder {

	if edit {
		return &InteractionCallbackBuilder{
			Type: 7,
		}
	} else {
		return &InteractionCallbackBuilder{
			Type: 4,
		}
	}
}

func CreateEmbed(title string, description string) *EmbedBuilder {

	return &EmbedBuilder{
		Title:       title,
		Description: description,
	}
}

func CreateEmbedVideo(url string) *EmbedBuilder {

	return &EmbedBuilder{
		Video: messages.Video{URL: url},
	}
}

func CreateAttachment(id int, description string, filename string) *AttachmentBuilder {
	return &AttachmentBuilder{
		ID:          id,
		Description: description,
		FileName:    filename,
	}
}

func CreateOption(label string, description string, value string) *OptionBuilder {

	return &OptionBuilder{
		Label:       label,
		Description: description,
		Value:       value,
		Emoji: messages.Emoji{
			ID:   "625891304148303894",
			Name: "rogue",
		},
	}
}

func CreateSelectComponent(label string, customID string) *ComponentBuilder {

	return &ComponentBuilder{
		Type:     3,
		Label:    label,
		CustomID: customID,
	}
}

func CreateActionRowComponent() *ActionRowComponentBuilder {

	return &ActionRowComponentBuilder{
		Type: 1,
	}
}

func CreateField(name string, value string) *FieldBuilder {

	return &FieldBuilder{
		Name:  name,
		Value: value,
	}
}

func (builder *InteractionCallbackBuilder) AddContent(content string) *InteractionCallbackBuilder {

	builder.Data = messages.Data{
		Content: content,
	}

	return builder
}

func (builder *InteractionCallbackBuilder) AddEmbed(embed *EmbedBuilder) *InteractionCallbackBuilder {

	builder.Data.Embeds = append(builder.Data.Embeds, *embed.Get())
	return builder
}

func (builder *InteractionCallbackBuilder) AddAttachment(attachment *AttachmentBuilder) *InteractionCallbackBuilder {

	builder.Data.Attachments = append(builder.Data.Attachments, *attachment.Get())
	return builder
}

func (builder *InteractionCallbackBuilder) AddActionRowComponent(component *ActionRowComponentBuilder) *InteractionCallbackBuilder {

	builder.Data.Components = append(builder.Data.Components, *component.Get())
	return builder
}

func (builder *ActionRowComponentBuilder) AddComponent(component *ComponentBuilder) *ActionRowComponentBuilder {

	builder.Components = append(builder.Components, *component.Get())
	return builder
}

func (builder *ComponentBuilder) AddOption(option *OptionBuilder) *ComponentBuilder {

	builder.Options = append(builder.Options, *option.Get())
	return builder
}

func (builder *EmbedBuilder) AddField(field *FieldBuilder) *EmbedBuilder {

	builder.Fields = append(builder.Fields, *field.Get())
	return builder
}

func (builder *InteractionCallbackBuilder) Get() *messages.InteractionCallback {

	return (*messages.InteractionCallback)(builder)
}

func (builder *EmbedBuilder) Get() *messages.Embed {

	return (*messages.Embed)(builder)
}

func (builder *AttachmentBuilder) Get() *messages.Attachment {

	return (*messages.Attachment)(builder)
}

func (builder *ActionRowComponentBuilder) Get() *messages.Components {

	return (*messages.Components)(builder)
}

func (builder *ComponentBuilder) Get() *messages.Component {

	return (*messages.Component)(builder)
}

func (builder *OptionBuilder) Get() *messages.Option {

	return (*messages.Option)(builder)
}

func (builder *FieldBuilder) Get() *messages.Field {

	return (*messages.Field)(builder)
}
